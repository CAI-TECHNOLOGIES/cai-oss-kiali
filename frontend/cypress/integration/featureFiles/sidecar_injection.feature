Feature: Controlling sidecar injection
    In Istio, at installation it is possible to set a default policy for automatic sidecar
    injection. In addition to the default policy, automatic sidecar injection can be
    controlled at namespace level and also at deployment level for more specific control.
    Kiali should provide the needed controls to override the default policy at namespace
    and deployment levels. Annotations are used to override the default policy and
    Kiali should reflect these annotations.

    Background:
        Given user is at administrator perspective

    Scenario: Override the default policy for automatic sidecar injection by enabling it in a namespace
        Given a namespace without override configuration for automatic sidecar injection
        When I override the default automatic sidecar injection policy in the namespace to enabled
        Then I should see the override annotation for sidecar injection in the namespace as "enabled"

    Scenario: Switch the override configuration for automatic sidecar injection in a namespace to disabled
        Given a namespace which has override configuration for automatic sidecar injection
        And the override configuration for sidecar injection is "enabled"
        When I change the override configuration for automatic sidecar injection policy in the namespace to "disable" it
        Then I should see the override annotation for sidecar injection in the namespace as "disabled"

    Scenario: Switch the override configuration for automatic sidecar injection in a namespace to enabled
        Given a namespace which has override configuration for automatic sidecar injection
        And the override configuration for sidecar injection is "disabled"
        When I change the override configuration for automatic sidecar injection policy in the namespace to "enable" it
        Then I should see the override annotation for sidecar injection in the namespace as "enabled"

    Scenario: Switch to using the default policy for automatic sidecar injection in a namespace
        Given a namespace which has override configuration for automatic sidecar injection
        When I remove override configuration for sidecar injection in the namespace
        Then I should see no override annotation for sidecar injection in the namespace

    Scenario: Override the default policy for automatic sidecar injection by enabling it in a workload
        Given a workload without a sidecar
        And the workload does not have override configuration for automatic sidecar injection
        When I override the default policy for automatic sidecar injection in the workload to "enable" it
        Then the workload should get a sidecar

    Scenario: Override the default policy for automatic sidecar injection by disabling it in a workload
        Given a workload with a sidecar
        And the workload does not have override configuration for automatic sidecar injection
        When I override the default policy for automatic sidecar injection in the workload to "disable" it
        Then the sidecar of the workload should vanish

    Scenario: Switch the override configuration for automatic sidecar injection in a workload to disabled
        Given a workload with a sidecar
        And the workload has override configuration for automatic sidecar injection
        When I change the override configuration for automatic sidecar injection in the workload to "disable" it
        Then the sidecar of the workload should vanish

    Scenario: Switch the override configuration for automatic sidecar injection in a workload to enabled
        Given a workload without a sidecar
        And the workload has override configuration for automatic sidecar injection
        When I change the override configuration for automatic sidecar injection in the workload to "enable" it
        Then the workload should get a sidecar

    Scenario: Remove override configuration for automatic sidecar injection in a workload
        Given a workload with override configuration for automatic sidecar injection
        When I remove override configuration for sidecar injection in the workload
        Then I should see no override annotation for sidecar injection in the workload
