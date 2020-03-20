// Copyright 2017 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package casbin

const (
	notImplemented = "not implemented"
)

// addPolicy adds a rule to the current policy.
func (e *Enforcer) addPolicy(sec string, ptype string, rule []string) (bool, error) {
	ruleAdded := e.model.AddPolicy(sec, ptype, rule)
	if !ruleAdded {
		return ruleAdded, nil
	}

	if e.adapter != nil && e.autoSave {
		if err := e.adapter.AddPolicy(sec, ptype, rule); err != nil {
			if err.Error() != notImplemented {
				return ruleAdded, err
			}
		}
	}

	if e.watcher !=nil && e.autoNotifyWatcher {
		err := e.watcher.Update()
		if err != nil {
			return ruleAdded, err
		}
	}

	return ruleAdded, nil
}

// addPolicies adds rules to the current policy.
// removePolicies removes rules from the current policy.
func (e *Enforcer) addPolicies(secs []string, ptypes []string, rules [][]string) ([]bool, []error) {
	rulesAdded := e.model.AddPolicies(secs, ptypes, rules)
	errs := make([]error, len(rulesAdded))
	for i, isRuleAdded := range rulesAdded {
		if !isRuleAdded {
			errs[i] = nil
		}
	}

	if e.adapter != nil && e.autoSave {
		newErrs := e.adapter.AddPolicies(secs, ptypes, rules)
		for i, err := range newErrs {
			if err != nil {
				if err.Error() != notImplemented {
					errs[i] = err
				}
			}
		}
		return rulesAdded, errs
	}

	if e.watcher !=nil && e.autoNotifyWatcher {
		err := e.watcher.Update()
		if err != nil {
			for i := range rulesAdded {
				errs[i] = err
			}
		}
	}

	return rulesAdded, errs
}

// removePolicy removes a rule from the current policy.
func (e *Enforcer) removePolicy(sec string, ptype string, rule []string) (bool, error) {
	ruleRemoved := e.model.RemovePolicy(sec, ptype, rule)
	if !ruleRemoved {
		return ruleRemoved, nil
	}

	if e.adapter != nil && e.autoSave {
		if err := e.adapter.RemovePolicy(sec, ptype, rule); err != nil {
			if err.Error() != notImplemented {
				return ruleRemoved, err
			}
		}
	}

	if e.watcher !=nil && e.autoNotifyWatcher {
		err := e.watcher.Update()
		if err != nil {
			return ruleRemoved, err
		}
	}

	return ruleRemoved, nil
}

// removePolicies removes rules from the current policy.
func (e *Enforcer) removePolicies(secs []string, ptypes []string, rules [][]string) ([]bool, []error) {
	rulesRemoved := e.model.RemovePolicies(secs, ptypes, rules)
	errs := make([]error, len(rulesRemoved))
	for i, isRuleRemoved := range rulesRemoved {
		if !isRuleRemoved {
			errs[i] = nil
		}
	}

	if e.adapter != nil && e.autoSave {
		newErrs := e.adapter.RemovePolicies(secs, ptypes, rules)
		for i, err := range newErrs {
			if err != nil {
				if err.Error() != notImplemented {
					errs[i] = err
				}
			}
		}
		return rulesRemoved, errs
	}

	if e.watcher !=nil && e.autoNotifyWatcher {
		err := e.watcher.Update()
		if err != nil {
			for i := range rulesRemoved {
				errs[i] = err
			}
		}
	}

	return rulesRemoved, errs
}

// removeFilteredPolicy removes rules based on field filters from the current policy.
func (e *Enforcer) removeFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) (bool, error) {
	ruleRemoved := e.model.RemoveFilteredPolicy(sec, ptype, fieldIndex, fieldValues...)
	if !ruleRemoved {
		return ruleRemoved, nil
	}

	if e.adapter != nil && e.autoSave {
		if err := e.adapter.RemoveFilteredPolicy(sec, ptype, fieldIndex, fieldValues...); err != nil {
			if err.Error() != notImplemented {
				return ruleRemoved, err
			}
		}
	}

	if e.watcher !=nil && e.autoNotifyWatcher {
		err := e.watcher.Update()
		if err != nil {
			return ruleRemoved, err
		}
	}

	return ruleRemoved, nil
}
