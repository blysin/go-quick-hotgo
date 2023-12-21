import { http, jumpExport } from '@/utils/http/axios';

// 获取供应商检索列表
export function List(params) {
  return http.request({
    url: '/supplier_search/vendor/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除供应商检索
export function Delete(params) {
  return http.request({
    url: '/supplier_search/vendor/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑供应商检索
export function Edit(params) {
  return http.request({
    url: '/supplier_search/vendor/save',
    method: 'POST',
    params,
  });
}

// 修改状态
export function ChangeStatus(params) {
  return http.request({
    url: '/supplier_search/vendor/change-status',
    method: 'POST',
    params,
  });
}


// 获取供应商检索指定详情
export function View(params) {
  return http.request({
    url: '/supplier_search/vendor/view',
    method: 'GET',
    params,
  });
}



// 导出供应商检索
export function Export(params) {
  jumpExport('/supplier_search/vendor/export', params);
}


// 获取供应商检索指定详情
export function GetCurrnecyList(params) {
  return http.request({
    url: '/supplier_search/currency/get',
    method: 'GET',
    params,
  });
}

// 获取供应商检索指定详情
export function SaveCurrnecyList(params) {
  return http.request({
    url: '/supplier_search/currency/save',
    method: 'POST',
    params,
  });
}


// 获取供应商检索指定详情
export function UploadFile(params) {
  let headers = {
    'Content-Type': 'multipart/form-data',
  }
  return http.request({
    url: '/supplier_search/ven-file/upload',
    method: 'POST',
    params,
    headers
  });
}


// 获取文件信息
export function GetVenFiles(params) {
  return http.request({
    url: '/supplier_search/ven-file/list',
    method: 'GET',
    params,
  });
}


// 获取文件信息
export function PageDetails(params) {
  return http.request({
    url: '/supplier_search/detail/list',
    method: 'GET',
    params,
  });
}

// 获取文件信息
export function Download(params) {
  return http.request({
    url: '/supplier_search/ven-file/download',
    method: 'GET',
    params,
  });
}

// 获取文件信息
export function GetVendor(params) {
  return http.request({
    url: '/supplier_search/vendor/get',
    method: 'GET',
    params,
  });
}
