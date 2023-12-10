# ApiHdfsSnapshotResult

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessedPathCount** | **float64** | Number of processed paths. | [optional] [default to null]
**ProcessedPaths** | **[]string** | The list of processed paths. &lt;p/&gt; This is only available in the full view. | [optional] [default to null]
**UnprocessedPathCount** | **float64** | Number of unprocessed paths. | [optional] [default to null]
**UnprocessedPaths** | **[]string** | The list of unprocessed paths. Note that paths that are currently being processed will also be included in this list. &lt;p/&gt; This is only available in the full view. | [optional] [default to null]
**CreatedSnapshotCount** | **float64** | Number of snapshots created. | [optional] [default to null]
**CreatedSnapshots** | [**[]ApiHdfsSnapshot**](ApiHdfsSnapshot.md) | List of snapshots created. &lt;p/&gt; This is only available in the full view. | [optional] [default to null]
**DeletedSnapshotCount** | **float64** | Number of snapshots deleted. | [optional] [default to null]
**DeletedSnapshots** | [**[]ApiHdfsSnapshot**](ApiHdfsSnapshot.md) | List of snapshots deleted. &lt;p/&gt; This is only available in the full view. | [optional] [default to null]
**CreationErrorCount** | **float64** | Number of errors detected when creating snapshots. | [optional] [default to null]
**CreationErrors** | [**[]ApiHdfsSnapshotError**](ApiHdfsSnapshotError.md) | List of errors encountered when creating snapshots. &lt;p/&gt; This is only available in the full view. | [optional] [default to null]
**DeletionErrorCount** | **float64** | Number of errors detected when deleting snapshots. | [optional] [default to null]
**DeletionErrors** | [**[]ApiHdfsSnapshotError**](ApiHdfsSnapshotError.md) | List of errors encountered when deleting snapshots. &lt;p/&gt; This is only available in the full view. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

