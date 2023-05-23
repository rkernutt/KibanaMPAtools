PUT _ingest/pipeline/metrics-system.diskio@custom
{
	"description": "Pipeline to create IOPS fields",
	"processors": [
		{
		  "set": {
			"field": "disk_read_iops",
			"value": "emit(doc['system.diskio.read.count'].value / (doc['system.diskio.read.time'].value / 1000))",
			"ignore_empty_value": true
		  }
		},
		{
		  "set": {
			"field": "disk_write_iops",
			"value": "emit(doc['system.diskio.write.count'].value / (doc['system.diskio.write.time'].value / 1000))",
			"ignore_empty_value": true
		  }
		}
	  ]
}