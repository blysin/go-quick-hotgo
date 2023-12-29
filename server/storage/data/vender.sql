create table hotgo.hg_fm_vendor
(
    id              int auto_increment comment '自增ID'
        primary key,
    vendor_name     varchar(32)                        not null comment '供应商名称',
    all_column      varchar(2048)                      not null comment '完整字段，多个用英文逗号隔开',
    required_column varchar(2048)                      not null comment '比填列，json格式',
    status          tinyint(2) default 0 not null comment '状态：0-新增，-1-删除，2已发布',
    created_at      datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    updated_at      datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    create_by       bigint null comment '创建人',
    update_by       bigint null comment '更新人',
    currency        varchar(8)                         not null comment '币种'
) comment '供应商主表';

create table hotgo.hg_fm_vendor_detail
(
    id                 int auto_increment comment '自增ID'
        primary key,
    vendor_id          int                                not null comment '供应商主表id',
    brand              varchar(32)                        not null comment '品牌',
    barcode            varchar(32)                        not null comment '条码',
    english_name       varchar(32)                        not null comment '英文名称',
    cost               int                                not null comment '成本、供货价',
    cost_cny           int                                not null comment '成本、供货价-人民币',
    selling_price      int                                not null comment '销售价格',
    selling_price_cny  int                                not null comment '销售价格-人民币',
    vendor             varchar(32)                        not null comment '供应商',
    currency           varchar(8)                         not null comment '币种',
    exchange_rate      decimal(10, 4)                     not null comment '汇率',
    exchange_rate_time datetime                           not null comment '汇率时间',
    vendor_data        text                               not null comment '工资内容,json格式存储',
    status             tinyint(2) default 0 not null comment '状态：0-新增，-1-删除，2已发布',
    created_at         datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    updated_at         datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    create_by          int                                not null comment '创建人',
    update_by          int                                not null comment '更新人'
) comment '供应商明细表';

create table hotgo.hg_fm_vendor_index
(
    id            int auto_increment comment '自增ID'
        primary key,
    unit_key      varchar(32)                        not null comment '唯一标识，当前已barcode作为标识',
    ven_detail_id int                                not null comment '关联的hg_fm_vendor_detail.id',
    status        tinyint(2) default 0 not null comment '状态：0-不可用，1-可用',
    created_at    datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    updated_at    datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间'
) comment '供应商索引表';

create table hotgo.hg_fm_vendor_upload_file
(
    id                     int auto_increment comment '自增ID'
        primary key,
    vendor_id              int null comment '供应商主表id',
    file_name              varchar(64)                        not null comment '文件名称',
    file_id                bigint                             not null comment '文件id',
    exception_data_file_id varchar(64) null comment '异常数据文件id',
    valid_num              int      default 0                 not null comment '正常数据条数',
    exception_num          int      default 0                 not null comment '异常数据条数',
    all_column             varchar(2048)                      not null comment '完整字段，多个用英文逗号隔开',
    created_at             datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    updated_at             datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    create_by              bigint null comment '创建人',
    update_by              bigint null comment '更新人'
) comment '供应商文件';

