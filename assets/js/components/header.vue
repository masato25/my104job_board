<template lang="pug">
    el-row(class="search-cond")
      el-col(:span=3)
        el-select(v-model="myicat" filterable placeholder="Select" @change="selectIcat")
          el-option(
              v-for="item in icats"
              :key="item[0]"
              :label="item[1]"
              :value="item[0]"
            )
      el-col(:span=6)
        el-input-number(@change="handleChangeSal" :min=0 :max=1000000 v-model="satmp")
</template>

<script>
import { mapState, mapActions } from 'vuex';

export default {
  name: 'AppHeader',
  data: () => {
    return {
      myicat: '',
      satmp: 0
    }
  },
  mounted: function() {
    this.myicat = this.$store.state.cat;
    this.satmp = this.$store.state.salarym;
  },
  methods: {
    selectIcat(myicat) {
      this.setIcat({icat: myicat})
    },
    handleChangeSal(satmp){
      this.setSalary({salary: satmp})
    },
    ...mapActions({
      setIcat: 'setIcat',
      setSalary: 'setSalary',
    })
  },
  computed: mapState(['salarym','icats']),
}
</script>
<style lang="scss">
.search-cond {
  .el-col {
    margin-right: 10px;
  }
}
</style>
