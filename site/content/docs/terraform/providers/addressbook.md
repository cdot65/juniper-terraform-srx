## 📌 Overview

The `cdot65.apstra.blueprint` module will allow you to manage the configuration of your Blueprint within Apstra.

Feature set as of version 0.0.15:

- manage Blueprints
- idempotent

## Example

Here is a basic example of using the module to manage your blueprint's configuration in Apstra

```yaml
{% raw %}
  - name: "### CREATE BLUEPRINT cicd_template"
    cdot65.apstra.blueprint:
        # define apstra server parameters
        server: "apstra.dmz.home"
        api_token: "{{ api_token }}"

        # define blueprint parameters
        design: "two_stage_l3clos"
        init_type: "template_reference"
        template_id: "{{ templates['data']['id'] }}"
        label: "cicd_template"

    # define whether to create or delete the blueprint
    state: present
{% endraw %}
```

## Options

If you'd like to see the options available for you within the module, have a look at the data model provided below.

| Option           | Type   | Description                                     |
| ---------------- | ------ | ----------------------------------------------- |
| `api_token`      | string | our API token to authenticate with Apstra       |
| `design`         | string | reference to our Design's UUID                  |
| `label`          | string | associate a label with our blueprint            |
| `port`           | int    | port number                                     |
| `server`         | string | Apstra's DNS hostname or IP address             |
| `state`          | string | determine whether to create or delete blueprint |
| `template_id`    | string | reference to our Template's UUID                |
| `validate_certs` | bool   | enable or disable SSL certificate validation    |

### Data Model

We can also get insight on all available options, and their expected input types, by looking at the module's data model.

```python

    @staticmethod
    def blueprint_spec():
        return dict(
            api_token=dict(
                required=True,
                fallback=(env_fallback, ['APSTRA_API_TOKEN', 'APSTRA_API_TOKEN', 'API_TOKEN']),
                no_log=True,
                type='str'
            ),
            design=dict(
                required=False,
                type='str'
            ),
            init_type=dict(
                required=False,
                type='str'
            ),
            label=dict(
                required=True,
                type='str'
            ),
            template_id=dict(
                required=False,
                type='str'
            ),
            port=dict(
                required=False,
                type='int'
            ),
            server=dict(
                required=False,
                type='str'
            ),
            state=dict(
                required=True,
                choices=['absent', 'present'],
                type='str'
            ),
            validate_certs=dict(
                type='bool',
                required=False,
                default=False
            ),
        )
```
