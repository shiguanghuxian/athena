package server

/* 本文件用于通过http上传数据到server */

// HTTPListener 开启一个http监听服务，用于通过http收集监控数据
func (ss *ServerService) HTTPListener() {
	// 判断是否开启了http收集数据的方式
	if ss.config.HTTP.Enable == false {
		return
	}
}
