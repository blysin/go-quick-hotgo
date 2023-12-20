import { h, ref } from 'vue';
import { NAvatar, NImage, NTag, NSwitch, NRate } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { Dicts } from '@/api/dict/dict';

import { isArray, isNullObject } from '@/utils/is';
import { getFileExt } from '@/utils/urlUtils';
import { defRangeShortcuts, defShortcuts, formatToDate } from '@/utils/dateUtil';
import { validate } from '@/utils/validateUtil';
import { getOptionLabel, getOptionTag, Options, errorImg } from '@/utils/hotgo';


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
  allColumn: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入完整字段，多个用英文逗号隔开',
  },
};

export const schemas = ref<FormSchema[]>([
  {
    field: 'id',
    component: 'NInputNumber',
    label: '自增ID',
    componentProps: {
      placeholder: '请输入自增ID',
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
]);

export const columns = [
  {
    title: '自增ID',
    key: 'id',
  },
  {
    title: '供应商名称',
    key: 'vendorName',
  },
  {
    title: '是否删除，0：未删除，1：已删除',
    key: 'isDeleted',
  },
  {
    title: '创建时间',
    key: 'createdAt',
  },
  {
    title: '更新时间',
    key: 'updatedAt',
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