curl -v -X POST -H "Content-Type: application/json" -d '{"source":{"bucket":"venicegeo-sample-data","key":"pointcloud/samp71-utm.laz"},"function":"info"}' http://192.168.99.100:8080/pdal

curl -v -X POST -H "Content-Type: application/json" -d '{"source":{"bucket":"venicegeo-sample-data","key":"pointcloud/samp71-utm.laz"},"function":"ground","destination":{"bucket":"venicegeo-sample-data","key":"temp/ground.laz"}}' http://192.168.99.100:8080/pdal

curl -v -X POST -H "Content-Type: application/json" -d '{"source":{"bucket":"venicegeo-sample-data","key":"pointcloud/samp71-utm.laz"},"function":"dtm","destination":{"bucket":"venicegeo-sample-data","key":"temp/dtm.tif"}}' http://192.168.99.100:8080/pdal

curl -v -X POST -H "Content-Type: application/json" -d '{"source":{"bucket":"venicegeo-sample-data","key":"pointcloud/samp71-utm.laz"},"function":"height","destination":{"bucket":"venicegeo-sample-data","key":"temp/height.laz"}}' http://192.168.99.100:8080/pdal

curl -v -X POST -H "Content-Type: application/json" -d '{"source":{"bucket":"venicegeo-sample-data","key":"pointcloud/samp71-utm.laz"},"function":"groundopts"}' http://192.168.99.100:8080/pdal

curl -v -X POST -H "Content-Type: application/json" -d '{"source":{"bucket":"venicegeo-sample-data","key":"pointcloud/samp71-utm.laz"},"fail":"info"}' http://192.168.99.100:8080/pdal

curl -v -X POST -H "Content-Type: application/json" -d '{"source":{"bucket":"venicegeo-sample-data","key":"pointcloud/samp71-utm.laz"},"function":"fail"}' http://192.168.99.100:8080/pdal

curl -v -X POST -H "Content-Type: application/json" -d '{"source":{"bucket":"bad-buckets","key":"pointcloud/samp71-utm.laz"},"function":"info"}' http://192.168.99.100:8080/pdal

curl -v -X POST -H "Content-Type: application/json" -d '{"source":{"bucket":"venicegeo-sample-data","key":"bad-folder/bad-file"},"function":"info"}' http://192.168.99.100:8080/pdal

curl -v -X POST -H "Content-Type: application/json" -d '{}' http://192.168.99.100:8080/pdal

curl -v -X POST http://192.168.99.100:8080/pdal

curl -v -X POST http://192.168.99.100:8080/ladp