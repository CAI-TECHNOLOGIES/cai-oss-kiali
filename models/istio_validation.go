package models

import (
	"encoding/json"

	"github.com/kiali/kiali/config"
	"github.com/kiali/kiali/log"
)

// NamespaceValidations represents a set of IstioValidations grouped by namespace
type NamespaceValidations map[string]IstioValidations

// IstioValidationKey is the key value composed of an Istio ObjectType and Name.
type IstioValidationKey struct {
	ObjectType string `json:"objectType"`
	Name       string `json:"name"`
	Namespace  string `json:"namespace"`
}

// IstioValidationSummary represents the number of errors/warnings of a set of Istio Validations.
type IstioValidationSummary struct {
	// Number of validations with error severity
	// required: true
	// example: 2
	Errors int `json:"errors"`
	// Number of Istio Objects analyzed
	// required: true
	// example: 6
	ObjectCount int `json:"objectCount"`
	// Number of validations with warning severity
	// required: true
	// example: 4
	Warnings int `json:"warnings"`
}

// IstioValidations represents a set of IstioValidation grouped by IstioValidationKey.
type IstioValidations map[IstioValidationKey]*IstioValidation

// IstioValidation represents a list of checks associated to an Istio object.
// swagger:model
type IstioValidation struct {
	// Name of the object itself
	// required: true
	// example: reviews
	Name string `json:"name"`

	// Type of the object
	// required: true
	// example: virtualservice
	ObjectType string `json:"objectType"`

	// Represents validity of the object: in case of warning, validity remains as true
	// required: true
	// example: false
	Valid bool `json:"valid"`

	// Array of checks. It might be empty.
	Checks []*IstioCheck `json:"checks"`

	// Related objects (only validation errors)
	References []IstioValidationKey `json:"references"`
}

// IstioCheck represents an individual check.
// swagger:model
type IstioCheck struct {
	// The check code used to identify a check
	// required: true
	// example: KIA0001
	Code string `json:"code"`

	// Description of the check
	// required: true
	// example: Weight sum should be 100
	Message string `json:"message"`

	// Indicates the level of importance: error or warning
	// required: true
	// example: error
	Severity SeverityLevel `json:"severity"`

	// String that describes where in the yaml file is the check located
	// example: spec/http[0]/route
	Path string `json:"path"`
}

type SeverityLevel string

const (
	ErrorSeverity   SeverityLevel = "error"
	WarningSeverity SeverityLevel = "warning"
	Unknown         SeverityLevel = "unknown"
)

var ObjectTypeSingular = map[string]string{
	"gateways":               "gateway",
	"virtualservices":        "virtualservice",
	"destinationrules":       "destinationrule",
	"serviceentries":         "serviceentry",
	"rules":                  "rule",
	"quotaspecs":             "quotaspec",
	"quotaspecbindings":      "quotaspecbinding",
	"policies":               "policy",
	"serviceroles":           "servicerole",
	"servicerolebindings":    "servicerolebinding",
	"clusterrbacconfigs":     "clusterrbacconfig",
	"authorizationpolicies":  "authorizationpolicy",
	"sidecars":               "sidecar",
	"peerauthentications":    "peerauthentication",
	"requestauthentications": "requestauthentication",
	"workloads":              "workload",
}

var checkDescriptors = map[string]IstioCheck{
	"authorizationpolicy.source.namespacenotfound": {
		Code:     "KIA0101",
		Message:  "Namespace not found for this rule",
		Severity: WarningSeverity,
	},
	"authorizationpolicy.source.principalnotfound": {
		Code:     "KIA0106",
		Message:  "Service Account not found for this principal",
		Severity: ErrorSeverity,
	},
	"authorizationpolicy.to.wrongmethod": {
		Code:     "KIA0102",
		Message:  "Only HTTP methods and fully-qualified gRPC names are allowed",
		Severity: WarningSeverity,
	},
	"authorizationpolicy.nodest.matchingregistry": {
		Code:     "KIA0104",
		Message:  "This host has no matching entry in the service registry",
		Severity: ErrorSeverity,
	},
	"authorizationpolicy.mtls.needstobeenabled": {
		Code:     "KIA0105",
		Message:  "This field requires mTLS to be enabled",
		Severity: ErrorSeverity,
	},
	"destinationrules.multimatch": {
		Code:     "KIA0201",
		Message:  "More than one DestinationRules for the same host subset combination",
		Severity: WarningSeverity,
	},
	"destinationrules.nodest.matchingregistry": {
		Code:     "KIA0202",
		Message:  "This host has no matching entry in the service registry (service, workload or service entries)",
		Severity: ErrorSeverity,
	},
	"destinationrules.nodest.subsetlabels": {
		Code:     "KIA0203",
		Message:  "This subset's labels are not found in any matching host",
		Severity: ErrorSeverity,
	},
	"destinationrules.trafficpolicy.notlssettings": {
		Code:     "KIA0204",
		Message:  "mTLS settings of a non-local Destination Rule are overridden",
		Severity: WarningSeverity,
	},
	"destinationrules.mtls.meshpolicymissing": {
		Code:     "KIA0205",
		Message:  "PeerAuthentication enabling mTLS at mesh level is missing",
		Severity: ErrorSeverity,
	},
	"destinationrules.mtls.nspolicymissing": {
		Code:     "KIA0206",
		Message:  "PeerAuthentication enabling namespace-wide mTLS is missing",
		Severity: ErrorSeverity,
	},
	"destinationrules.mtls.policymtlsenabled": {
		Code:     "KIA0207",
		Message:  "PeerAuthentication with TLS strict mode found, it should be permissive",
		Severity: ErrorSeverity,
	},
	"destinationrules.mtls.meshpolicymtlsenabled": {
		Code:     "KIA0208",
		Message:  "PeerAuthentication enabling mTLS found, permissive mode needed",
		Severity: ErrorSeverity,
	},
	"destinationrules.nodest.subsetnolabels": {
		Code:     "KIA0209",
		Message:  "This subset has not labels",
		Severity: WarningSeverity,
	},
	"gateways.multimatch": {
		Code:     "KIA0301",
		Message:  "More than one Gateway for the same host port combination",
		Severity: WarningSeverity,
	},
	"gateways.selector": {
		Code:     "KIA0302",
		Message:  "No matching workload found for gateway selector in this namespace",
		Severity: WarningSeverity,
	},
	"generic.exportto.namespacenotfound": {
		Code:     "KIA0005",
		Message:  "No matching namespace found or namespace is not accessible",
		Severity: ErrorSeverity,
	},
	"generic.multimatch.selectorless": {
		Code:     "KIA0002",
		Message:  "More than one selector-less object in the same namespace",
		Severity: ErrorSeverity,
	},
	"generic.multimatch.selector": {
		Code:     "KIA0003",
		Message:  "More than one object applied to the same workload",
		Severity: ErrorSeverity,
	},
	"generic.selector.workloadnotfound": {
		Code:     "KIA0004",
		Message:  "No matching workload found for the selector in this namespace",
		Severity: WarningSeverity,
	},
	"peerauthentication.mtls.destinationrulemissing": {
		Code:     "KIA0401",
		Message:  "Mesh-wide Destination Rule enabling mTLS is missing",
		Severity: ErrorSeverity,
	},
	"peerauthentications.mtls.destinationrulemissing": {
		Code:     "KIA0501",
		Message:  "Destination Rule enabling namespace-wide mTLS is missing",
		Severity: ErrorSeverity,
	},
	"peerauthentications.mtls.disabledestinationrulemissing": {
		Code:     "KIA0505",
		Message:  "Destination Rule disabling namespace-wide mTLS is missing",
		Severity: ErrorSeverity,
	},
	"peerauthentications.mtls.disablemeshdestinationrulemissing": {
		Code:     "KIA0506",
		Message:  "Destination Rule disabling mesh-wide mTLS is missing",
		Severity: ErrorSeverity,
	},
	"port.appprotocol.mismatch": {
		Code:     "KIA0602",
		Message:  "Port appProtocol must follow <protocol> form",
		Severity: ErrorSeverity,
	},
	"port.name.mismatch": {
		Code:     "KIA0601",
		Message:  "Port name must follow <protocol>[-suffix] form",
		Severity: ErrorSeverity,
	},
	"service.deployment.port.mismatch": {
		Code:     "KIA0701",
		Message:  "Deployment exposing same port as Service not found",
		Severity: WarningSeverity,
	},
	"serviceentries.workloadentries.addressmatch": {
		Code:     "KIA1201",
		Message:  "Missing one or more addresses from matching WorkloadEntries",
		Severity: WarningSeverity,
	},
	"servicerole.invalid.services": {
		Code:     "KIA0901",
		Message:  "Unable to find all the defined services",
		Severity: ErrorSeverity,
	},
	"servicerole.invalid.namespace": {
		Code:     "KIA0902",
		Message:  "ServiceRole can only point to current namespace",
		Severity: ErrorSeverity,
	},
	"servicerolebinding.invalid.role": {
		Code:     "KIA0903",
		Message:  "ServiceRole does not exists in this namespace",
		Severity: ErrorSeverity,
	},
	"sidecar.egress.servicenotfound": {
		Code:     "KIA1004",
		Message:  "This host has no matching entry in the service registry",
		Severity: WarningSeverity,
	},
	"sidecar.global.selector": {
		Code:     "KIA1006",
		Message:  "Global default sidecar should not have workloadSelector",
		Severity: WarningSeverity,
	},
	"virtualservices.gateway.oldnomenclature": {
		Code:     "KIA1108",
		Message:  "Preferred nomenclature: <gateway namespace>/<gateway name>",
		Severity: Unknown,
	},
	"virtualservices.nohost.hostnotfound": {
		Code:     "KIA1101",
		Message:  "DestinationWeight on route doesn't have a valid service (host not found)",
		Severity: ErrorSeverity,
	},
	"virtualservices.nogateway": {
		Code:     "KIA1102",
		Message:  "VirtualService is pointing to a non-existent gateway",
		Severity: ErrorSeverity,
	},
	"virtualservices.route.singleweight": {
		Code:     "KIA1104",
		Message:  "The weight is assumed to be 100 because there is only one route destination",
		Severity: WarningSeverity,
	},
	"virtualservices.route.repeatedsubset": {
		Code:     "KIA1105",
		Message:  "This host subset combination is already referenced in another route destination",
		Severity: WarningSeverity,
	},
	"virtualservices.singlehost": {
		Code:     "KIA1106",
		Message:  "More than one Virtual Service for same host",
		Severity: WarningSeverity,
	},
	"virtualservices.subsetpresent.subsetnotfound": {
		Code:     "KIA1107",
		Message:  "Subset not found",
		Severity: WarningSeverity,
	},
	"validation.unable.cross-namespace": {
		Code:     "KIA0001",
		Message:  "Unable to verify the validity, cross-namespace validation is not supported for this field",
		Severity: Unknown,
	},
	"workload.authorizationpolicy.needstobecovered": {
		Code:     "KIA1201",
		Message:  "This workload is not covered by any authorization policy",
		Severity: WarningSeverity,
	},
}

func Build(checkId string, path string) IstioCheck {
	check := checkDescriptors[checkId]
	check.Path = path
	return check
}

func BuildKey(objectType, name, namespace string) IstioValidationKey {
	return IstioValidationKey{ObjectType: objectType, Namespace: namespace, Name: name}
}

func CheckMessage(checkId string) string {
	if val, ok := checkDescriptors[checkId]; ok {
		return val.GetFullMessage()
	} else {
		return "ISTIO CHECK ID DOES NOT EXIST:" + checkId
	}
}

func (ic IstioCheck) GetFullMessage() string {
	return ic.Code + " " + ic.Message
}

func (iv IstioValidations) FilterBySingleType(objectType, name string) IstioValidations {
	fiv := IstioValidations{}
	for k, v := range iv {
		// We don't want to filter other types
		if k.ObjectType != objectType {
			fiv[k] = v
		} else {
			// But for this exact type we're strict
			if k.Name == name {
				fiv[k] = v
			}
		}
	}

	return fiv
}

func (iv IstioValidations) FilterByKey(objectType, name string) IstioValidations {
	fiv := IstioValidations{}
	for k, v := range iv {
		if k.Name == name && k.ObjectType == objectType {
			fiv[k] = v
		}
	}

	return fiv
}

// FilterByTypes takes an input as ObjectTypes, transforms to singular types and filters the validations
func (iv IstioValidations) FilterByTypes(objectTypes []string) IstioValidations {
	types := make(map[string]bool, len(objectTypes))
	for _, objectType := range objectTypes {
		types[ObjectTypeSingular[objectType]] = true
	}
	fiv := IstioValidations{}
	for k, v := range iv {
		if _, found := types[k.ObjectType]; found {
			fiv[k] = v
		}
	}

	return fiv
}

func (iv IstioValidations) MergeValidations(validations IstioValidations) IstioValidations {
	for key, validation := range validations {
		v, ok := iv[key]
		if !ok {
			iv[key] = validation
		} else {
		AddUnique:
			for _, toAdd := range validation.Checks {
				for _, existing := range v.Checks {
					if toAdd.Path == existing.Path &&
						toAdd.Severity == existing.Severity &&
						toAdd.Message == existing.Message {
						continue AddUnique
					}
				}
				v.Checks = append(v.Checks, toAdd)
			}
			v.Valid = v.Valid && validation.Valid
		AddUniqueReference:
			for _, toAdd := range validation.References {
				for _, existing := range v.References {
					if toAdd == existing {
						continue AddUniqueReference
					}
				}
				v.References = append(v.References, toAdd)
			}
		}
	}
	return iv
}

func (iv IstioValidations) MergeReferences(validations IstioValidations) IstioValidations {
	for _, currentValidations := range iv {
		if currentValidations.References == nil {
			currentValidations.References = make([]IstioValidationKey, 0, len(validations))
		}
		for k := range validations {
			currentValidations.References = append(currentValidations.References, k)
		}
	}

	return iv
}

func (iv IstioValidations) SummarizeValidation(ns string) IstioValidationSummary {
	ivs := IstioValidationSummary{}
	for k, v := range iv {
		if k.Namespace == ns {
			ivs.mergeSummaries(v.Checks)
		}
	}
	return ivs
}

func (summary *IstioValidationSummary) mergeSummaries(cs []*IstioCheck) {
	for _, c := range cs {
		if c.Severity == ErrorSeverity {
			summary.Errors += 1
		} else if c.Severity == WarningSeverity {
			summary.Warnings += 1
		}
	}
	summary.ObjectCount += 1
}

// MarshalJSON implements the json.Marshaler interface.
func (iv IstioValidations) MarshalJSON() ([]byte, error) {
	out := make(map[string]map[string]*IstioValidation)
	for k, v := range iv {
		_, ok := out[k.ObjectType]
		if !ok {
			out[k.ObjectType] = make(map[string]*IstioValidation)
		}
		out[k.ObjectType][k.Name+"."+k.Namespace] = v
	}
	return json.Marshal(out)
}

func (iv *IstioValidations) StripIgnoredChecks() {
	// strip away codes that are to be ignored
	codesToIgnore := config.Get().KialiFeatureFlags.Validations.Ignore
	if len(codesToIgnore) > 0 {
		for curValidationKey, curValidation := range *iv {
			idx := 0
			// loop over each IstioCheck in the current Validation and only keep it if it is not ignored
			for _, curCheck := range curValidation.Checks {
				ignoreCheck := false
				for _, cti := range codesToIgnore {
					if cti == curCheck.Code {
						ignoreCheck = true
						log.Tracef("Ignoring validation failure [%+v] for object [%s:%s] in namespace [%s]", curCheck, curValidationKey.ObjectType, curValidationKey.Name, curValidationKey.Namespace)
						break
					}
				}
				if !ignoreCheck {
					curValidation.Checks[idx] = curCheck
					idx++
				}
			}
			// Prevent memory leak - nil out ignored checks
			for extraIdx := idx; extraIdx < len(curValidation.Checks); extraIdx++ {
				curValidation.Checks[extraIdx] = nil
			}
			curValidation.Checks = curValidation.Checks[:idx]
		}
	}
}
