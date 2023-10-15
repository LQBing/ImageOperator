# ImageOperator

A tool for operate many images.

examples:

```shell
imageoperator pull origin

imageoperator pull mirror

imageoperator pull pyarmor

imageoperator tag origin private

imageoperator tag mirror private

imageoperator tag pyarmor private

imageoperator push private

imageoperator save private
```


Usage:

```shell
imageoperator config_help
imageoperator [--dry-run] [-c <config_file_path>] pull <registry_group_type>
imageoperator [--dry-run] [-c <config_file_path>] push <registry_group_type>
imageoperator [--dry-run] [-c <config_file_path>] rmi <registry_group_type>
imageoperator [--dry-run] [-c <config_file_path>] list <registry_group_type>
imageoperator [--dry-run] [-c <config_file_path>] tag <registry_group_type> <registry_group_type>
imageoperator [--dry-run] [-c <config_file_path>] [-o <package_path>] save <registry_group_type>
imageoperator [--dry-run] [-c <config_file_path>] [-i <package_path>] load <registry_group_type>
```

Json config instruction:

- Default config file is `image.json`

- `.registry_groups[<registry_group_type>]` for storage registry_group registry and group

- `.images[i][<registry_groups>]` for storage iamges in registry_group

- `.images[i][<registry_groups>]["image"]` If you mean `mysql`, do not fill it with `nginx`, please fill it with `docker.io/library/mysql`. Or this image will be combine with `.registry_groups[<registry_group_type>]["registry"]` and `.registry_groups[<registry_group_type>]["group"]`

Base structure:

```json
{
	"registry_groups": {
		"origin": {
			"registry": "your.registry.com",
			"group": "yourgroup"
		}
	},
	"images": [{
		"origin": {
			"image": "docker.io/library/redis",
			"tag": "alpine3.18"
		},
		"private": {
			"image": "mirrors",
			"tag": "library_redis_alpine3.18"
		}
	}]
}
```

Example image.json:

```json
{
	"registry_groups": {
		"origin": {
			"registry": "your.registry.com",
			"group": "yourgroup"
		},
		"private": {
			"registry": "dockerhub.kubekey.local",
			"group": "yourgroup"
		}
	},
	"images": [{
			"origin": {
				"image": "tools",
				"tag": "alpine"
			},
			"private": {
				"image": "tools",
				"tag": "alpine"
			}
		},
		{
			"origin": {
				"image": "docker.io/library/redis",
				"tag": "alpine3.18"
			},
			"private": {
				"image": "mirrors",
				"tag": "library_redis_alpine3.18"
			}
		},
		{
			"origin": {
				"image": "docker.io/library/mysql",
				"tag": "8.0"
			},
			"private": {
				"image": "mirrors",
				"tag": "library_mysql_8.0"
			}
		}
	]
}
```
