<template>
  <div>
    <div class="n-layout-page-header">
      <!--      <n-card :bordered="false" title="供应商检索">-->
      <!--          这是由系统生成的CURD表格，你可以将此行注释改为表格的描述-->
      <!--      </n-card>-->
    </div>
    <n-card :bordered="false" class="proCard">

      <BasicForm
        @register="register"
        @submit="reloadTable"
        @reset="reloadTable"
        @keyup.enter="reloadTable"
        ref="searchFormRef"
      >
        <template #statusSlot="{ model, field }">
          <n-input v-model:value="model[field]"/>
        </template>
      </BasicForm>

      <BasicTable
        :openChecked="true"
        :columns="columns"
        :request="loadDataTable"
        :row-key="(row) => row.id"
        ref="actionRef"
        :actionColumn="actionColumn"
        :checked-row-keys="checkedIds"
        @update:checked-row-keys="onCheckedRow"
        :scroll-x="1090"
        :resizeHeightOffset="-10000"
        size="small"
      >
        <template #tableTitle>
          <n-button
            type="primary"
            @click="addTable"
            class="min-left-space"
            v-if="hasPermission(['/supplier_search/vendor/edit'])"
          >
            <template #icon>
              <n-icon>
                <PlusOutlined/>
              </n-icon>
            </template>
            添加
          </n-button>
          <n-button
            type="error"
            @click="handleBatchDelete"
            :disabled="batchDeleteDisabled"
            class="min-left-space"
            v-if="hasPermission(['/supplier_search/vendor/delete'])"
          >
            <template #icon>
              <n-icon>
                <DeleteOutlined/>
              </n-icon>
            </template>
            批量删除
          </n-button>
<!--          <n-button-->
<!--            type="primary"-->
<!--            @click="managerCurrency"-->
<!--            class="min-left-space"-->
<!--            v-if="hasPermission(['/supplier_search/vendor/export'])"-->
<!--          >-->
<!--            <template #icon>-->
<!--              <n-icon>-->
<!--                <ExportOutlined/>-->
<!--              </n-icon>-->
<!--            </template>-->
<!--            导出-->
<!--          </n-button>-->

          <n-button
            type="primary"
            @click="managerCurrency"
            class="min-left-space"
          >
            <template #icon>
              <n-icon>
                <DollarOutlined/>
              </n-icon>
            </template>
            币种管理
          </n-button>

        </template>
      </BasicTable>
    </n-card>
    <Edit
      @reloadTable="reloadTable"
      @updateShowModal="updateShowModal"
      :showModal="showModal"
      :formParams="formParams"
      :fullCurrency="fullCurrency"
    />
    <EditCurrency
      @reloadTable="reloadTable"
      @updateShowCurrencyModal="updateShowCurrencyModal"
      :showModal="showCurrencyModal"
      :formParams="formParams"
    />
  </div>
</template>

<script lang="ts" setup>
import {h, onMounted, reactive, ref} from 'vue';
import {useDialog, useMessage} from 'naive-ui';
import {BasicTable, TableAction} from '@/components/Table';
import {BasicForm, useForm} from '@/components/Form/index';
import {usePermission} from '@/hooks/web/usePermission';
import {
  ChangeStatus,
  Delete,
  Export,
  GetCurrnecyList,
  List
} from '@/api/addons/supplier_search/vendor';
import {columns, Currency, newState, schemas, State, Status, StatusList} from './model';
import {DeleteOutlined, DollarOutlined, ExportOutlined, PlusOutlined} from '@vicons/antd';
import {useRouter} from 'vue-router';
import Edit from './edit.vue';
import EditCurrency from './edit_currency.vue';

const {hasPermission} = usePermission();
const router = useRouter();
const actionRef = ref();
const dialog = useDialog();
const message = useMessage();
const searchFormRef = ref<any>({});
const batchDeleteDisabled = ref(true);
const checkedIds = ref([]);
const showModal = ref(false);
const showCurrencyModal = ref(false);
const formParams = ref<State>();
const fullCurrency = ref<Currency[]>([]);

const actionColumn = reactive({
  width: 300,
  title: '操作',
  key: 'action',
  // fixed: 'right',
  render(record) {
    let btns = {
      style: 'button',
      actions: [
      ],
      // dropDownActions: [
      //   {
      //     label: '查看详情',
      //     key: 'view',
      //     auth: ['/supplier_search/vendor/view'],
      //   },
      // ],
      // select: (key) => {
      //   if (key === 'view') {
      //     return handleView(record);
      //   }
      // },
    };
    let size = 'tiny';
    let delBtn = {
      label: '删除',
      size,
      onClick: handleDelete.bind(null, record),
      auth: ['/supplier_search/vendor/delete'],
    };
    let editBtn = {
      label: '编辑',
      size,
      onClick: handleEdit.bind(null, record),
      auth: ['/supplier_search/vendor/edit'],
    };
    let recBtn = {
      label: '还原',
      size,
      type: 'warning',
      onClick: handleRestore.bind(null, record),
      auth: ['/supplier_search/vendor/delete'],
    };
    let pubBtn = {
      label: '发布',
      size,
      type: 'success',
      onClick: handlePublish.bind(null, record),
      auth: ['/supplier_search/vendor/delete'],
    };
    let backBtn = {
      label: '撤回',
      size,
      type: 'warning',
      onClick: handleBack.bind(null, record),
      auth: ['/supplier_search/vendor/delete'],
    };

    if (record.status === Status.normal.value) {
      btns.actions.push(editBtn);
      btns.actions.push(delBtn);
      btns.actions.push(pubBtn);
    }
    if (record.status === Status.published.value) {
      btns.actions.push(backBtn);
    }
    if (record.status === Status.delete.value) {
      btns.actions.push(recBtn);
    }
    return h(TableAction as any,btns );
  },
});

const [register, {}] = useForm({
  gridProps: {cols: '2xl:4 s:1 m:2 l:3 xl:4 2xl:4'},
  labelWidth: 80,
  schemas,
});

const loadDataTable = async (res) => {
  let page = await List({...searchFormRef.value?.formModel, ...res});

  let statusMap = StatusList.reduce((map, obj) => {
    map[obj.value] = obj.label;
    return map;
  }, {});

  page.list.forEach((item) => {
    item.statusName = statusMap[item.status];
  });
  return page
};

function addTable() {
  showModal.value = true;
  formParams.value = newState(null);
}

function updateShowModal(value) {
  showModal.value = value;
  reloadTable();
}

function updateShowCurrencyModal(value) {
  loadCurrency();
  showCurrencyModal.value = value;
}


function onCheckedRow(rowKeys) {
  batchDeleteDisabled.value = rowKeys.length <= 0;
  checkedIds.value = rowKeys;
}

function reloadTable() {
  actionRef.value.reload();
}

function handleView(record: Recordable) {
  router.push({name: 'vendorView', params: {id: record.id}});
}

function handleEdit(record: Recordable) {
  showModal.value = true;
  formParams.value = newState(record as State);
}

function handleDelete(record: Recordable) {
  dialog.warning({
    title: '警告',
    content: '确定要删除？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      Delete(record).then((_res) => {
        message.success('删除成功');
        reloadTable();
      });
    },
    onNegativeClick: () => {
      // message.error('取消');
    },
  });
}

function handleRestore(record: Recordable) {
  dialog.warning({
    title: '请确认',
    content: '确定要还原？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      let params = {
        "vendor_id": record.id,
        "detail_id": null,
        "status": Status.normal.value
      }
      ChangeStatus(params).then((_res) => {
        message.success('发布成功');
        reloadTable();
      });
    },
    onNegativeClick: () => {
      // message.error('取消');
    },
  });
}

function handleBack(record: Recordable) {
  dialog.warning({
    title: '请确认',
    content: '确定要撤回？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      let params = {
        "vendor_id": record.id,
        "detail_id": null,
        "status": Status.normal.value
      }
      ChangeStatus(params).then((_res) => {
        message.success('发布成功');
        reloadTable();
      });
    },
    onNegativeClick: () => {
      // message.error('取消');
    },
  });
}

function handlePublish(record: Recordable) {
  dialog.success({
    title: '请确认',
    content: '确定要发布？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      let params = {
        "vendor_id": record.id,
        "detail_id": null,
        "status": Status.published.value
      }
      ChangeStatus(params).then((_res) => {
        message.success('发布成功');
        reloadTable();
      });
    },
    onNegativeClick: () => {
      // message.error('取消');
    },
  });
}

function handleBatchDelete() {
  dialog.warning({
    title: '警告',
    content: '确定要批量删除？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      Delete({id: checkedIds.value}).then((_res) => {
        batchDeleteDisabled.value = true;
        checkedIds.value = [];
        message.success('删除成功');
        reloadTable();
      });
    },
    onNegativeClick: () => {
      // message.error('取消');
    },
  });
}

function loadCurrency() {
  GetCurrnecyList({}).then((res) => {
    if (res && res.list) {
      res.list.forEach((item) => {
        item.desc = item.desc + '（' + item.name + '）';
      });
    }

    fullCurrency.value = res.list
  });
}

function handleExport() {
  message.loading('正在导出列表...', {duration: 1200});
  Export(searchFormRef.value?.formModel);
}

function managerCurrency() {
  // message.info('币种管理');
  showCurrencyModal.value = true;
}

onMounted(async () => {
  loadCurrency();
});

</script>

<style lang="less" scoped></style>
