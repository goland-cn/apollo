package initialize

import (
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
)

func ApolloInit(appid, cluster, namespace, secret string, IsBackupConfig bool) (agollo.Client, error) {
	// 创建一个Apollo配置对象
	c := &config.AppConfig{
		AppID:          appid,                         // 替换为您的AppID
		Cluster:        cluster,                       // Apollo集群名称，默认为"default"
		IP:             "http://106.14.141.61:31080/", // 替换为您的Apollo配置中心的地址和端口
		NamespaceName:  namespace,                     // 替换为您的应用程序所使用的命名空间名称
		IsBackupConfig: IsBackupConfig,                // 是否备份配置
		Secret:         secret,                        // 替换为您的秘钥
	}
	// 使用Apollo配置对象初始化Agollo客户端
	client, err := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})
	return client, err
}
