<template>
  <div>
    <n-spin :show="loading" description="请稍候...">
      <n-modal
        v-model:show="isShowModal"
        :show-icon="false"
        :style="{
          width: dialogWidth,
        }"
        :title="'编辑币种'"
        preset="dialog"
      >
        <n-scrollbar style="max-height: 300px">
          <n-table :bordered="false" :single-line="true" size="small">
            <thead>
            <tr>
              <th>编码</th>
              <th>名称</th>
              <th>操作</th>
            </tr>
            </thead>

            <tbody>
            <tr v-for="(item,index) in currencyList">
              <td>{{ item.name }}</td>
              <td>{{ item.desc }}</td>
              <td>
                <n-icon v-if="index!==0" class="arrow-btn" size="30" title="上移"
                        @click="move(item.name,-1)">
                  <ArrowUp/>
                </n-icon>
                <n-icon v-if="index !== currencyList.length -1" class="arrow-btn" size="30"
                        title="下移" @click="move(item.name,1)">
                  <ArrowDown/>
                </n-icon>

              </td>
            </tr>
            </tbody>

          </n-table>
        </n-scrollbar>
        <template #action>
          <n-space>
            <n-button @click="closeForm">取消</n-button>
            <n-button :loading="formBtnLoading" type="info" @click="confirmForm">确定</n-button>
          </n-space>
        </template>
      </n-modal>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
import {computed, onMounted, ref, watch} from 'vue';
import {Currency, newState, State} from './model';
import {useMessage} from 'naive-ui';
import {adaModalWidth} from '@/utils/hotgo';
import {ArrowDown, ArrowUp} from '@vicons/ionicons5'
import {GetCurrnecyList, SaveCurrnecyList} from "@/api/addons/supplier_search/vendor";

const emit = defineEmits(['reloadTable', 'updateShowCurrencyModal']);

interface Props {
  showModal: boolean;
  formParams?: State;
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
    emit('updateShowCurrencyModal', value);
  },
});

const loading = ref(false);
const params = ref<State>(props.formParams);
const message = useMessage();
const formRef = ref<any>({});
const dialogWidth = ref('75%');
const formBtnLoading = ref(false);
const currencyList = ref<Currency[]>([]);


onMounted(async () => {
  adaModalWidth(dialogWidth);
});

function closeForm() {
  isShowModal.value = false;
}

function loadForm(value) {
  loading.value = true;
  GetCurrnecyList({}).then((res) => {
    currencyList.value = res.list
    loading.value = false;
  });
}

function move(name, direction) {
  const index = currencyList.value.findIndex((item) => item.name === name);
  const item = currencyList.value[index];
  currencyList.value.splice(index, 1);
  currencyList.value.splice(index + direction, 0, item);
}

function confirmForm() {
  formBtnLoading.value = true;
  let params = {
    list: currencyList.value
  }
  SaveCurrnecyList(params).then((res) => {
    message.success('操作成功');
    formBtnLoading.value = false;
    closeForm();
  });
}

watch(
  () => props.showModal,
  (value) => {
    if (value) {
      loadForm(value);
    }
  }
);
</script>

<style lang="less">
.arrow-btn {
  cursor: pointer;
}

</style>
