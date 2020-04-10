# cfPing
 Find the best cloudflare ip to your location.
# 配置
 在ini文件中有详细说明。
# 过程
 1.向cfip.txt文件中的ip进行本地ping测试，找到按丢包最少（第一要素）延迟最低（第二要素）进行筛选。
 2.利用dnspod api，自动更改dns记录。
# 有效性
 旨在找到对自己本地最好的节点。
