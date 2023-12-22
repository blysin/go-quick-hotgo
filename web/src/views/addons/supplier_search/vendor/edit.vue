<template>
  <div>
    <n-spin :show="loading" description="请稍候...">
      <n-modal
        v-model:show="isShowModal"
        :show-icon="false"
        preset="dialog"
        :title="params?.id > 0 ? '编辑供应商 #' + params?.id : '添加供应商数据'"
        :style="{
          width: dialogWidth,
        }"
      >
        <n-steps
          size="small"
          :current="(current as number)"
          :status="currentStatus"
          v-show="showStep"
        >
          <n-step
            title="上传数据"
            description="仅支持上传.xlsx数据"
          />
          <n-step
            title="修改基础信息"
            description="修改数据基础信息"
          />
          <n-step
            title="发布"
            description="保存完数据后需要发布才能生效"
          />
          <n-step
            title="成功"
            description="数据创建完成"
          />
        </n-steps>

        <div v-if="current === 1">

          <Upload :venId="0" @uploadSuccessFunc="uploadSuccess"/>

        </div>

        <n-form
          v-if="current === 2 || current === 3"
          :model="params"
          :rules="rules"
          ref="formRef"
          size="small"
          label-placement="left"
          :label-width="100"
          class="py-4"
        >
          <n-divider title-placement="left" v-show="showStep">
            基础信息
          </n-divider>
          <n-form-item label="供应商名称" path="vendorName">
            <n-input placeholder="请输入供应商名称" v-model:value="params.vendorName"/>
          </n-form-item>

          <n-form-item label="币种" path="exchange">
            <n-select
              label-field="desc"
              value-field="name"
              v-model:value="params.exchange"
              filterable
              placeholder="请选择币种"
              :options="fullCurrency"
            />
          </n-form-item>

          <n-divider title-placement="left">
            字段确认
          </n-divider>

          <n-grid x-gap="12" :cols="3">
            <n-gi>
              <n-form-item
                label="品牌名字段"
                path="presetColumn.brandName">
                <n-select
                  v-model:value="params.presetColumn.brandName"
                  filterable
                  placeholder="选择字段"
                  :options="filterFullColumns"
                />
              </n-form-item>
            </n-gi>
            <n-gi>
              <n-form-item label="英文名字段"
                           path="presetColumn.enName">
                <n-select
                  v-model:value="params.presetColumn.enName"
                  filterable
                  placeholder="选择字段"
                  :options="filterFullColumns"
                />
              </n-form-item>
            </n-gi>
            <n-gi>
              <n-form-item label="品牌商字段"
                           path="presetColumn.vendorName">
                <n-select
                  v-model:value="params.presetColumn.vendorName"
                  filterable
                  placeholder="选择字段"
                  :options="filterFullColumns"
                />
              </n-form-item>
            </n-gi>
          </n-grid>

          <n-grid x-gap="12" :cols="3">
            <n-gi>
              <n-form-item label="销售价字段"
                           path="presetColumn.salePrice">
                <n-select
                  v-model:value="params.presetColumn.salePrice"
                  filterable
                  placeholder="选择字段"
                  :options="filterFullColumns"
                />
              </n-form-item>
            </n-gi>
            <n-gi>
              <n-form-item label="供货价字段"
                           path="presetColumn.supplyPrice">
                <n-select
                  v-model:value="params.presetColumn.supplyPrice"
                  filterable
                  placeholder="选择字段"
                  :options="filterFullColumns"
                />
              </n-form-item>
            </n-gi>
            <n-gi>
              <n-form-item label="条码字段"
                           path="presetColumn.barCode">
                <n-select
                  v-model:value="params.presetColumn.barCode"
                  filterable
                  placeholder="选择字段"
                  :options="filterFullColumns"
                />
              </n-form-item>
            </n-gi>
          </n-grid>


        </n-form>


        <template #action>
          <n-space>
            <n-button @click="closeForm">取消</n-button>
            <n-button type="info" :loading="formBtnLoading" @click="nextStep"
                      :disabled="!uploadResponse || !uploadResponse.id"
                      v-if="current === 1">下一步
            </n-button>
            <n-button strong secondary type="warning" :loading="formBtnLoading" @click="preStep"
                      v-if="current !== 1 && params.id === 0">重新上传文件
            </n-button>
            <n-button type="info" :loading="formBtnLoading" @click="publish"
                      v-if="params.id !== 0 && params.status === 0">发布
            </n-button>
            <n-button type="info" :loading="formBtnLoading" @click="confirmForm"
                      v-if="current === 2 || current === 3">确定
            </n-button>
          </n-space>
        </template>

      </n-modal>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
import {computed, onMounted, ref, watch} from 'vue';
import {ChangeStatus, Edit, GetVendor} from '@/api/addons/supplier_search/vendor';
import {Currency, newState, rules, SaveParam, State, Status, UploadResponse} from './model';
import {StepsProps, useMessage} from 'naive-ui';
import {adaModalWidth} from '@/utils/hotgo';
import Upload from './upload.vue';

// 入参
const emit = defineEmits(['reloadTable', 'updateShowModal']);

interface Props {
  showModal: boolean;
  formParams?: State;
  fullCurrency?: Currency[];
}

const props = withDefaults(defineProps<Props>(), {
  showModal: false,
  formParams: () => {
    return newState(null);
  },
});

const isShowModal = computed({
  get: () => {
    return props.showModal;
  },
  set: (value) => {
    emit('updateShowModal', value);
  },
});

// 局部变量
const loading = ref(false);
const params = ref<SaveParam>({
  id: 0,
  vendorName: '',
  fileId: 0,
  status: 0,
  exchange: '',
  presetColumn: {
    brandName: '',
    barCode: '',
    enName: '',
    supplyPrice: '',
    salePrice: '',
    vendorName: '',
  },
});
const message = useMessage();
const formRef = ref<any>({});
const dialogWidth = ref('75%');
const formBtnLoading = ref(false);
const current = ref<number | null>(1);
const currentStatus = ref<StepsProps['status']>('process');
const fullColumns = ref<any[]>([]);
const uploadResponse = ref<UploadResponse>({
  brandName: '',
  barCode: '',
  enName: '',
  supplyPrice: '',
  salePrice: '',
  vendorName: '',
  id: 0,
  file_name: '',
  all_columns: []
});
const showStep = ref(false);

// 计算属性，用于过滤fullColumns
const filterFullColumns = computed(() => {
  let options = [];
  fullColumns.value.forEach((item) => {
    let opt = {
      label: item.label,
      value: item.value,
    };
    if (item.value === params.value.presetColumn.brandName ||
      item.value === params.value.presetColumn.barCode ||
      item.value === params.value.presetColumn.enName ||
      item.value === params.value.presetColumn.supplyPrice ||
      item.value === params.value.presetColumn.salePrice ||
      item.value === params.value.presetColumn.vendorName) {
      opt.disabled = true;
    }


    options.push(opt);
  });
  return options;
});


function nextStep(e) {
  e && e.preventDefault();
  if (current.value === 1) {
    current.value = 2;
    currentStatus.value = 'process';
  } else if (current.value === 2) {
    current.value = 3;
    currentStatus.value = 'process';
  } else if (current.value === 3) {
    current.value = 4;
    currentStatus.value = 'process';
  }
}

function preStep(e) {
  e && e.preventDefault();
  if (current.value === 0) current.value = null
  else if (current.value === null) current.value = 4
  else current.value--
}

function uploadSuccess(e) {
  uploadResponse.value = e;
  message.success('上传成功，请填写基础信息');

  if (e.file_name) {
    params.value.vendorName = e.file_name;
  }

  if (e.all_columns) {
    let options = [];
    e.all_columns.forEach((item) => {
      let opt = {
        label: item,
        value: item,
      };
      options.push(opt);
    });
    fullColumns.value = options;
  }

  if (e.brandName) {
    params.value.presetColumn.brandName = e.brandName;
  }
  if (e.barCode) {
    params.value.presetColumn.barCode = e.barCode;
  }
  if (e.enName) {
    params.value.presetColumn.enName = e.enName;
  }
  if (e.supplyPrice) {
    params.value.presetColumn.supplyPrice = e.supplyPrice;
  }
  if (e.salePrice) {
    params.value.presetColumn.salePrice = e.salePrice;
  }
  if (e.vendorName) {
    params.value.presetColumn.vendorName = e.vendorName;
  }

  params.value.fileId = e.id;

  nextStep(null);
}

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      console.log("保存数据params", params.value)
      Edit(params.value).then((_res) => {
        message.success('操作成功', _res);

        if (params.value.id === 0) {
          nextStep(null);
        }
        params.value.id = _res.id;
      });
    } else {
      message.error('请填写完整信息');
    }
    formBtnLoading.value = false;
  });
}

function publish(e) {
  e.preventDefault();
  let data = {
    "vendor_id": params.value.id,
    "detail_id": null,
    "status": Status.published.value
  }
  ChangeStatus(data).then((_res) => {
    message.success('发布成功', _res);
    closeForm();
    emit('reloadTable');
  });
}

onMounted(async () => {
  adaModalWidth(dialogWidth);
});

function closeForm() {
  isShowModal.value = false;
}

function loadForm(value) {
  let uploadData = {
    brandName: '',
    barCode: '',
    enName: '',
    supplyPrice: '',
    salePrice: '',
    vendorName: '',
    id: 0,
    file_name: '',
    all_columns: []
  }
  let data = {
    id: 0,
    vendorName: '',
    fileId: 0,
    status: 0,
    exchange: '',
    presetColumn: {
      brandName: '',
      barCode: '',
      enName: '',
      supplyPrice: '',
      salePrice: '',
      vendorName: '',
    },
  };

  console.log("loadForm", value);
  if (value.id) {
    current.value = 2;
    showStep.value = false;
    GetVendor({
      id: value.id
    }).then((_res) => {
      console.log("编辑", _res);
      data.id = _res.id;
      data.vendorName = _res.vendorName;
      data.status = _res.status;
      data.exchange = _res.currency;

      data.presetColumn = JSON.parse(_res.requiredColumn);

      let all_columns = JSON.parse(_res.allColumn);
      if (all_columns) {
        let options = [];
        all_columns.forEach((item) => {
          let opt = {
            label: item,
            value: item,
          };
          options.push(opt);
        });
        fullColumns.value = options;
      }

      params.value = data;

      console.log("编辑", data, uploadData);
      loading.value = false;
    });

  } else {
    current.value = 1;
    params.value = data;
    uploadResponse.value = uploadData;
    showStep.value = true;
  }




}

watch(
  () => props.formParams, (value) => {
    loadForm(value);
  }
);
</script>

<style lang="less">
</style>
