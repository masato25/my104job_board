<template lang="pug">
el-menu(:default-active="activeIndex" class="el-menu-nav")
  el-menu-item(index="1") HelloJbo104
    i(class="el-icon-cold-drink")
</template>

<script>
export default {
  data: () => {
    return {}
  }
}
</script>

<style lang="scss">
.el-menu {
  background-color: #545c64;
  .el-menu-item {
    color: #ecd611;
    font-weight: 600;
  }
}
</style>
