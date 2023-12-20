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
    url: '/supplier_search/vendor/edit',
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