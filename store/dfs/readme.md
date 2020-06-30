

# dfs(distributed-file-system)

- https://github.com/gostor/awesome-go-storage#file-system
- https://github.com/topics/distributed-file-system?l=go


## ref:

- https://github.com/pkwenda/Blog/issues/18
- https://juejin.im/post/5ddd4af66fb9a07189523d79
- http://blog.zollty.com/b/archive/comparison-of-distributed-file-storage-MinIO-SeaweedFS-FastDFS.html


## projects:


### minio

- https://github.com/minio/minio
    - golang, 国人开发项目
    - 易用性更高
- https://docs.min.io/cn/
- https://zhuanlan.zhihu.com/p/103803549
- 兼容 s3
- sdk: https://github.com/minio/minio-go
- MinIO是兼容Amazon S3的，换句话说，MinIO可以伪装成Amazon S3，你可以用Amazon S3的SDK操作MinIO。
- MinIO支持多租户，但是却不支持动态扩展。因此，大租户，就单独搭一套MinIO吧。小租户倒是可以共用一套。
- https://github.com/pkwenda/Blog/issues/18  

#### 配合 thumbor:

- https://github.com/thumbor/thumbor
- 搭建分布式图片存储
- https://thumbor.readthedocs.io/en/latest/index.html


### SeaweedFS:

- https://github.com/chrislusf/seaweedfs
- 适合存储小文件
- golang
- [图片服务器minio+thumbor](https://jamesdeng.github.io/2018/11/13/%E5%9B%BE%E7%89%87%E6%9C%8D%E5%8A%A1%E5%99%A8Minio+Thumbor.html)

### go-fastdfs:

- https://github.com/sjqzhang/go-fastdfs
- docs: https://sjqzhang.github.io/go-fastdfs/#character
- web ui: https://github.com/perfree/go-fastdfs-web
- golang




### HDFS:

- 适合存储大文件

### etc:

- https://github.com/douban/gobeansdb
    - douban
    - kv, golang
- https://github.com/storj/storj
    - golang
- https://github.com/chubaofs/chubaofs
    - golang
- https://github.com/gluster/glusterfs
- https://github.com/ceph/ceph
- https://github.com/moosefs/moosefs
- https://docs.mongodb.com/manual/core/gridfs/


## libs:

- https://github.com/spf13/afero


