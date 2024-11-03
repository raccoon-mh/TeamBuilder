<template>

  <!-- Page header -->
  <div class="page-header d-print-none">
    <div class="container-xl">
      <div class="row g-2 align-items-center">
        <div class="col">
          <h2 class="page-title">
            Admin Page
          </h2>
        </div>
      </div>
    </div>
  </div>

  <!-- Page body -->

  <div class="page-body">
    <div class="container-xl">
      <div class="card">
        <div class="card-header">
          <h3 class="card-title">Most Visited Pages</h3>
        </div>
        <div class="card-table table-responsive">
          <div id="example-table"></div>
        </div>
      </div>
    </div>
  </div>

</template>

<script>
import { apiGet } from '../util/http';
import { tableInit } from '../util/table';


export default {
  name: 'AdminPage',
  async mounted() {
    const columns = [
      { title: "ID", field: "ID", width: 150 },
      { title: "IP", field: "IP" },
      { title: "Count", field: "Count" },
      { title: "CreatedAt", field: "CreatedAt", sorter: "date", hozAlign: "center" },
      { title: "UpdatedAt", field: "UpdatedAt", sorter: "date", hozAlign: "center" },
      { title: "DeletedAt", field: "DeletedAt" },
    ];
    tableInit(await this.fetchVisitor(), "example-table", columns);
  },
  methods: {
    async fetchVisitor() {
      try {
        const visitor = await apiGet("/admin/visitor/list");
        if (visitor.status == 200) {
          console.log('user:', visitor.data.res);
          return visitor.data.res;
        } else {
          console.error('error fetchUserInfo:', visitor.status);
          return [];
        }
      } catch (error) {
        console.error('error fetchUserInfo:', error);
        return [];
      }
    },
  }
}
</script>
