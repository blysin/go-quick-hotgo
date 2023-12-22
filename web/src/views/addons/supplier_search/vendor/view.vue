<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="供应商检索详情"> <!-- CURD详情页--> </n-card>
    </div>
    <n-card :bordered="false" class="proCard mt-4" size="small" :segmented="{ content: true }">
      <n-descriptions label-placement="left" class="py-2" column="3">
        <n-descriptions-item>
          <template #label>供应商名称</template>
          {{ formValue.vendorName }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>创建时间</template>
          {{ formValue.createdAt }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>创建人</template>
          {{ formValue.createBy }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>币种</template>
          {{ formValue.currency }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>状态</template>
          <n-tag :bordered="false" :type="statusTag">
            {{ getStatus(formValue.status) }}
          </n-tag>
        </n-descriptions-item>

      </n-descriptions>

      <n-tabs type="line" animated>
        <n-tab-pane name="明细列表" tab="明细列表">
          <n-space>
            <n-form
              label-placement="left"
              label-width="auto"
              require-mark-placement="right-hanging"
              size="small"
              inline
              :style="{
                maxWidth: '800px'
              }"
            >
              <n-form-item label="品牌：" path="brand">
                <n-input placeholder="" @keyup.enter="formQuery" v-bind:on-clear="formQuery"
                         clearable v-model:value="formParam.brand"/>
              </n-form-item>
              <n-form-item label="条码：" path="barcode">
                <n-input placeholder="" @keyup.enter="formQuery" v-bind:on-clear="formQuery"
                         clearable v-model:value="formParam.barcode"/>
              </n-form-item>
              <n-form-item label="供应商：" path="vendor">
                <n-input placeholder="" @keyup.enter="formQuery" v-bind:on-clear="formQuery"
                         clearable v-model:value="formParam.vendor"/>
              </n-form-item>
              <n-form-item label="状态：" path="status">
                <select v-model="formParam.status" @change="formQuery"
                        class="cus-select n-input n-input--resizable n-input--stateful">
                  <option value="-99">全部</option>
                  <option v-for="item in statusList" :value="item.value">{{ item.label }}</option>
                </select>
              </n-form-item>
            </n-form>

            <n-button type="primary" size="small" @click="formQuery">查询</n-button>
          </n-space>

          <n-spin :show="isLoading">
            <n-table :bordered="false" :single-line="false" size="small">
              <thead>
              <tr>
                <th></th>
                <th>品牌</th>
                <th>英文名称</th>
                <th>条码</th>
                <th>供应商</th>
                <th>供货价</th>
                <th>供货价(人民币)</th>
                <th>销售价格</th>
                <th>销售价格(人民币)</th>
                <th>汇率</th>
                <th>状态</th>
                <th>操作</th>
              </tr>
              </thead>
              <tbody>
              <tr v-for="(item,i) in detailList">
                <td>{{ i + 1 }}</td>
                <td>{{ item.brand }}</td>
                <td>{{ item.englishName }}</td>
                <td>{{ item.barcode }}</td>
                <td>{{ item.vendor }}</td>
                <td>{{ splitIntoYuan(item.cost) }}</td>
                <td>{{ splitIntoYuan(item.costCny) }}</td>
                <td>{{ splitIntoYuan(item.sellingPrice) }}</td>
                <td>{{ splitIntoYuan(item.sellingPriceCny) }}</td>
                <td v-bind:title="'汇率时间：'+item.exchangeRateTime">{{ item.exchangeRate }}</td>
                <td>{{ getStatus(item.status) }}</td>
                <td>
                  <n-space>
                    <n-button type="primary" size="tiny" @click="viewData(item)">查看</n-button>

                    <n-popconfirm
                      v-if="item.status === Status.published.value"
                      @positive-click="changeStatus(item.id,Status.normal)"
                    >
                      <template #trigger>
                        <n-button type="error" size="tiny">撤销</n-button>
                      </template>
                      是否确认撤销该数据？
                    </n-popconfirm>

                    <n-popconfirm
                      v-if="item.status === Status.normal.value"
                      @positive-click="changeStatus(item.id,Status.published)"
                    >
                      <template #trigger>
                        <n-button type="success" size="tiny">发布</n-button>
                      </template>
                      是否确认发布该数据？
                    </n-popconfirm>
                  </n-space>
                </td>
              </tr>
              </tbody>
            </n-table>
            <div class="pagination">
              <n-pagination
                class="pagination"
                v-model:page="page.current"
                v-model:page-size="page.pageSize"
                v-model:item-count="page.total"
                v-bind:on-update:page="changePage"
                v-bind:on-update:page-size="changePageSize"

                show-size-picker
                :page-sizes="[10, 20, 30, 40]"
              >
                <template #suffix="{ itemCount }">
                  共 {{ itemCount }} 条
                </template>
              </n-pagination>
            </div>
          </n-spin>


        </n-tab-pane>
        <n-tab-pane name="文件列表" tab="原始文件">
          <n-table :bordered="false" :single-line="false" size="small">
            <thead>
            <tr>
              <th>文件名称</th>
              <th>数据量</th>
              <th>上传时间</th>
              <th>上传用户</th>
            </tr>
            </thead>
            <tbody>
            <tr v-for="item in formValue.files">
              <td>
                <div class="down-item" @click="downloadFile(item.id)">{{ item.fileName }}</div>
              </td>
              <td>{{ item.validNum }}</td>
              <td>{{ item.createdAt }}</td>
              <td>{{ item.createBy }}</td>
            </tr>
            </tbody>
          </n-table>
        </n-tab-pane>
      </n-tabs>
    </n-card>

    <n-modal v-model:show="showModal" :title="currentItem.brand +'详情'" positive-text="确认"
             @positive-click="showModal = false">
      <n-card
        style="width: 600px"
        :title="currentItem.brand +'详情'"
        :bordered="false"
        role="dialog"
        aria-modal="true"
      >
        <template #header-extra>
          <n-button text style="font-size: 24px;color: black;" @click="showModal = false">
            <n-icon>
              <cash-icon/>
            </n-icon>
          </n-button>
        </template>

        <n-scrollbar style="max-height: 500px">
          <n-table :bordered="false" :single-line="false" striped size="medium">
            <tbody>
            <tr v-for="item in currentItem.dataArray">
              <td style="text-align: right;width: 10em;">{{ item.name }}:</td>
              <td>{{ item.value }}</td>
            </tr>

            </tbody>
          </n-table>
        </n-scrollbar>
        <br/>

        <template #footer>
        </template>
      </n-card>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
import {computed, onMounted, ref} from 'vue';
import {useRouter} from 'vue-router';
import {FormInst, useMessage} from 'naive-ui';
import {ChangeStatus, Download, GetVendor, PageDetails} from '@/api/addons/supplier_search/vendor';
import {newState, Status, StatusList} from './model';
import {CloseOutline as CashIcon} from '@vicons/ionicons5'

const message = useMessage();
const router = useRouter();
const id = Number(router.currentRoute.value.params.id);
const formValue = ref(newState(null));
const fileAvatarCSS = computed(() => {
  return {
    '--n-merged-size': `var(--n-avatar-size-override, 80px)`,
    '--n-font-size': `18px`,
  };
});

const showFiles = ref(false);
const page = ref({
  current: 1,
  pageSize: 10,
  total: 0,
});
const statusList = ref<any>([])
const detailList = ref<any>([])
const isLoading = ref(true)
const formParam = ref({
  brand: '',
  barcode: '',
  vendor: '',
  status: -99,
})
const formRef = ref<FormInst | null>(null)
const showModal = ref(false)
const currentItem = ref({})

function downloadFile(fileId) {
  Download({id: fileId}).then(res => {
    const url = res.url
    message.info('开始下载文件' + url);
    download(url);
  });
}

//下载
async function download(url: string) {
  url = location.protocol + "//" + location.host + "/" + url;
  window.open(url)

  // 创建隐藏的可下载链接
  // const eleLink = document.createElement('a');
  // eleLink.download = '';
  // eleLink.style.display = 'none';
  // 字符内容转变成blob地址
  // console.log("url",url)
  // let bytes = await http.request({
  //   url: url,
  //   method: 'GET',
  // });
  //
  // const blob = new Blob([bytes]);
  // eleLink.href = URL.createObjectURL(blob);
  // // 触发点击
  // document.body.appendChild(eleLink);
  // eleLink.click();
  // // 然后移除
  // document.body.removeChild(eleLink);
}

function changePage(_page) {
  console.log("分页变动", _page)
  if (page.value.current === _page) {
    return;
  }
  page.value.current = _page
  fetchData()
}

function changePageSize(_size) {
  console.log("分页变动", _size)
  if (page.value.pageSize === _size) {
    return;
  }
  page.value.current = 1
  page.value.pageSize = _size
  fetchData()
}

function formQuery() {
  console.log("表单参数", formParam.value)
  page.value.current = 1
  fetchData()
}

function changeStatus(itemId, status) {
  console.log('变更状态', itemId, status);
  let data = {
    "vendor_id": id,
    "detail_id": itemId,
    "status": status.value
  }
  ChangeStatus(data).then((_res) => {
    message.success('操作成功', _res);
    fetchData();
  });
}

function viewData(item) {
  console.log("查看数据", item)

  let data = {};
  if (item.vendorData) {
    data = JSON.parse(item.vendorData)
  }

  let array = [];

  for (let key in data) {
    let item = {
      name: key,
      value: data[key]
    }
    array.push(item)
  }

  item.dataArray = array

  currentItem.value = item;
  showModal.value = true;
}

function fetchData() {
  isLoading.value = true
  let params = {
    page: page.value.current,
    pageSize: page.value.pageSize,
    vendorId: id,
    ...formParam.value
  }
  PageDetails(params).then((res) => {
    detailList.value = res.list
    page.value.total = res.totalCount
  }).finally(() => isLoading.value = false)
}

function splitIntoYuan(val) {
  return val && (val / 100).toFixed(2)
}

function getStatus(val) {
  let status = StatusList.filter(item => item.value === val)
  return status?.length > 0 ? status[0].label : val
}

//计算属性
const statusTag = computed(() => {
  let tag = ""
  let status = formValue.value.status || 0

  if (Status.delete.value === status) {
    tag = "error";
  } else if (Status.published.value === status) {
    tag = "success";
  }

  return tag
});

onMounted(async () => {
  if (!id || id < 1) {
    message.error('自增ID不正确，请检查！');
    return;
  }

  formValue.value = await GetVendor({id: id})

  // StatusList 去掉value = -1 的项
  let status = StatusList.filter(item => item.value !== -1)

  statusList.value = status
  fetchData()
});
</script>

<style lang="less" scoped>
.down-item {
  cursor: pointer;
  color: #1890ff;

  &:hover {
    color: #40a9ff;
  }
}

// 分页，靠右
.pagination {
  margin-top: 5px;
  float: right;
}

.cus-select {
  width: 200px;
  border: 1px solid #e0e0e6;
  border-radius: 3px;
  height: 30px;
}

.desc-label {
  color: red;
}
</style>
