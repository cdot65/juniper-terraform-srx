## 📌 Overview

The `cdot65.apstra.resources` module will allow you to manage the configuration of your Blueprint within Apstra.

Feature set as of version 0.0.7:

- manage all resources
- supports tagging
- idempotent

Suported Resources:

| Resource Type | Ansible module parameter "type" |
| ------------- | ------------------------------- |
| ASN Pool      | asn-pools                       |
| IPv4 Pool     | ip-pools                        |
| IPv6 Pool     | ipv6-pools                      |
| VLAN Pool     | vlan-pools                      |
| VNI Pool      | vni-pools                       |

## Example

Here is a basic example of using the module to manage your resources within in Apstra

```yaml
{% raw %}
  - name: Create an IP Pool Resource with two prefixes
    cdot65.apstra.resources:
        # define server connectivity options
        server: apstra.dmz.home
        port: 443
        validate_certs: False
        api_token: "{{ api_token }}"

        # define resource allocations
        display_name: "Windows Servers"
        tags: []
        type: "ip-pools"
        subnets:
            - "100.1.1.0/24"
            - "100.1.2.0/24"

        # state whether you want to create or delete this resource
        state: present

    # store the output of our task as a new variable to debug later
    register: windows_servers_ippool

{% endraw %}
```

## Options

If you'd like to see the options available for you within the module, have a look at the data model provided below.

| Option           | Type   | Description                                    |
| ---------------- | ------ | ---------------------------------------------- |
| `api_token`      | string | our API token to authenticate with Apstra      |
| `address`        | string | standby                                        |
| `asn`            | int    | standby                                        |
| `display_name`   | string | standby                                        |
| `ipv6_address`   | string | standby                                        |
| `port`           | int    | standby                                        |
| `ranges`         | list   | standby                                        |
| `subnets`        | list   | standby                                        |
| `server`         | string | Apstra's DNS hostname or IP address            |
| `state`          | string | determine whether to create or delete resource |
| `tags`           | list   | standby                                        |
| `type`           | string | standby                                        |
| `validate_certs` | bool   | enable or disable SSL certificate validation   |

### Data Model

We can also get insight on all available options, and their expected input types, by looking at the module's data model.

```python

    @staticmethod
    def resources_spec():
        return dict(
            address=dict(
                required=False,
                type='str'),
            asn=dict(
                required=False,
                type='int'),
            api_token=dict(
                required=True,
                fallback=(env_fallback, ['APSTRA_API_TOKEN', 'APSTRA_API_TOKEN', 'API_TOKEN']),
                no_log=True,
                type='str'),
            display_name=dict(
                required=True,
                fallback=(env_fallback, ['APSTRA_USERNAME', 'APSTRA_USERNAME', 'USERNAME']),
                type='str'),
            ipv6_address=dict(
                required=False,
                type='str'),
            port=dict(
                required=True,
                type='int'),
            ranges=dict(
                required=False,
                type='list',
                elements='dict',
                options=dict(
                    first=dict(
                        required=True,
                        type='int'),
                    last=dict(
                        required=True,
                        type='int'),
                    ),
                ),
            server=dict(
                required=False,
                type='str'),
            state=dict(
                required=False,
                choices=['absent', 'present'],
                type='str'),
            subnets=dict(
                required=False,
                type='list',
                elements='str'),
            tags=dict(
                required=False,
                type='list',
                elements='str'),
            type=dict(
                required=True,
                choices=['asn-pools', 'external-routers', 'ip-pools', 'ipv6-pools', 'vlan-pools', 'vni-pools'],
                type='str'),
            validate_certs=dict(
                type='bool',
                required=False,
                default=False),
        )
```
