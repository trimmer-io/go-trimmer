## v1.3 [2018-08-04]

  * New API version 2018-08-04
  * new asset states: `published`, `reviewing`, `approved`, `rejected`, `archived`
  * new `AssetStatistics` fields: versions, views, downloads
  * asset version support
    - new `Asset` fields: `version`, `locked`, `parentId`
    - new fork type 'version-forks'
    - new fields IN `AssetForkParams`: `version`, `locked`
    - new fields in `AssetListParams`: `version`, `original`, `head`
    - new fields in `LinkListParams`: `version`, `original`, `head`

### Breaking API changes

- replaced asset states `analyzing` and `transcoding` with more generic state `processing`
- renamed asset field `version` (hash of metadata revision) to `revision`
- added new asset field `version` (user-defined string, default HEAD)
- renamed asset metadata type from `MetaVersion` to `MetaRevision`
- renamed asset metadata id field from `versionId` to `revisionId`
- renamed asset metadata (hash) field `version` to `revision`
- renamed asset metadata query result list from `versions` to `revisions` (this is internal and should not be noticable)

## v1.0.2 [2018-07-04]

  * New API version 2018-07-04
  * Smart stash support and stash description
  * Add commit messages to metadata updates
  * Add job state group `all` to list all jobs
  * Extended VFX media families
  * Extended list of media formats

## v1.0.1 [2018-04-11]

  * Supports new API version 2018-04-10

## v1.0.0 [2017-11-13]

  * Initial Release
  * Supporting API version 2017-08-04
