<template>
  <n-upload
    multiple
    directory-dnd
    accept=".xlsx"
    action="/supplier_search/ven-file/upload"
    :max="1"
    :custom-request="customRequest"
    :data="{
      id: props.venId,
    }"
  >
    <n-upload-dragger>
      <div style="margin-bottom: 12px">
        <n-icon size="48" :depth="3">
          <archive-icon/>
        </n-icon>
      </div>
      <n-text style="font-size: 16px">
        点击或者拖动文件到该区域来上传
      </n-text>
      <n-p depth="3" style="margin: 8px 0 0 0">
        仅支持上传.xlsx数据
      </n-p>
    </n-upload-dragger>
  </n-upload>
</template>
<script lang="ts" setup>
import {UploadCustomRequestOptions, useMessage} from 'naive-ui'
import {UploadFile} from "@/api/addons/supplier_search/vendor";

//基础组件
const message = useMessage()

// 入参，1、主表id，非必填；2、上传成功回调函数
const emit = defineEmits(['uploadSuccessFunc']);
const props = defineProps<{
  venId?: number;
}>();

// 局部变量

// 方法
const customRequest = ({
                         file,
                         data,
                         headers,
                         withCredentials,
                         action,
                         onFinish,
                         onError,
                         onProgress
                       }: UploadCustomRequestOptions) => {
  const formData = new FormData()
  if (data) {
    Object.keys(data).forEach((key) => {
      formData.append(
        key,
        data[key as keyof UploadCustomRequestOptions['data']]
      )
    })
  }
  formData.append("file", file.file as File)

  UploadFile(formData).then((resp) => {
    console.log("upload resp",resp)
    emit('uploadSuccessFunc', resp)
    onFinish()
  }).catch((error) => {
    message.success(error.message)
    onError()
  })
}


</script>
<style scoped lang="scss">
</style>
