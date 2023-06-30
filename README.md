# changelog_yaml


Example yaml file: changelog.yaml

```
versioning_format: "[major.minor.patch] - [build] - date"
versions:
  - version: "1.0.1"
    build: "2" 
    date: "2023-06-30"
    changes:    
      added:  
        - "[RPC] example rpc"
        - "something here"       
      fixed: 
        - "[Notifications] Wrong name in notification when sending message"       
  - version: "1.0.0"
    build: "1"
    date: "2023-06-30"
    changes:
      added:
        - "Entire code base, initial release"
```

output: changelog.md
```
# Changelog

versioning format: `[major.minor.patch] - [build] - date`

## [1.0.1] - [2] - 2023-06-30

### Added

- [RPC] example rpc
- something here

### Fixed

- [Notifications] Wrong name in notification when sending message


## [1.0.0] - [1] - 2023-06-30

### Added

- Entire code base, initial release
```
