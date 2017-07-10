package plugins

import (
	_ "github.com/luopengift/transport/plugins/file"
	_ "github.com/luopengift/transport/plugins/hdfs"
	_ "github.com/luopengift/transport/plugins/kafka"
	_ "github.com/luopengift/transport/plugins/elasticsearch"
	_ "github.com/luopengift/transport/plugins/std"
)

func init() {}
