import {ref} from 'vue';
import {cloneDeep} from 'lodash-es';
import {FormSchema} from '@/components/Form';
import {defRangeShortcuts} from '@/utils/dateUtil';


export interface State {
  id: number;
  vendorName: string;
  allColumn: string;
  requiredColumn: string;
  isDeleted: number;
  createdAt: string;
  updatedAt: string;
  createBy: number;
  updateBy: number;
}

export const defaultState: State = {
  id: 0,
  vendorName: '',
  allColumn: '',
  requiredColumn: '',
  isDeleted: 0,
  createdAt: '',
  updatedAt: '',
  createBy: 0,
  updateBy: 0,
};

export function newState(state: State | null): State {
  if (state !== null) {
    return cloneDeep(state);
  }
  return cloneDeep(defaultState);
}


export const rules = {
  vendorName: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入供应商名称',
  },
  exchange: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请选择',
  },

  "presetColumn.brandName": {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请选择',
  },
  "presetColumn.barCode": {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请选择',
  },
  "presetColumn.enName": {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请选择',
  },
  "presetColumn.supplyPrice": {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请选择',
  },
  "presetColumn.salePrice": {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请选择',
  },
  "presetColumn.vendorName": {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请选择',
  },

};

export const Status = {
  normal: {
    value: 0,
    label: "未发布"
  },
  delete: {
    value: -1,
    label: "已删除"
  },
  published: {
    value: 2,
    label: "已发布"
  },
}

export const StatusList = [
  {
    value: 0,
    label: "未发布"
  },
  {
    value: -1,
    label: "已删除"
  },
  {
    value: 2,
    label: "已发布"
  },
]

export const schemas = ref<FormSchema[]>([
  {
    field: 'vendorName',
    component: 'NInput',
    label: '名称',
    componentProps: {
      placeholder: '请输入',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'createdAt',
    component: 'NDatePicker',
    label: '创建时间',
    componentProps: {
      type: 'datetimerange',
      clearable: true,
      shortcuts: defRangeShortcuts(),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'status',
    component: 'NSelect',
    label: '状态',
    componentProps: {
      placeholder: '选择',
      options: StatusList,
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

export const columns = [
  {
    title: '名称',
    key: 'vendorName',
  },
  {
    title: '币种',
    key: 'currency',
  },
  {
    title: '状态',
    key: 'statusName',
  },
  {
    title: '创建时间',
    key: 'createdAt',
  },
  {
    title: '创建人',
    key: 'createBy',
  },
  {
    title: '更新人',
    key: 'updateBy',
  },
];

export interface UploadResponse {
  brandName: string;
  barCode: string;
  enName: string;
  supplyPrice: string;
  salePrice: string;
  vendorName: string;
  id: number;
  file_name: string;
  all_columns: string[];
}

export interface PresetColumn {
  brandName: string;
  barCode: string;
  enName: string;
  supplyPrice: string;
  salePrice: string;
  vendorName: string;
}

export interface SaveParam {
  id: number;
  status: number;
  vendorName: string;
  fileId: number;
  exchange: string;
  presetColumn: PresetColumn;
}


export interface Currency {
  name: string;
  desc: string;
}

