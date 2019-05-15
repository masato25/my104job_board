const mutations = {
  SET_JOBS_LIST(state, { list }){
    state.jobs = list.data
    state.totalpage = list.pagging.total_entries_size
    state.currentPage = list.pagging.page
  },
  SET_PAGE_Limit(state, { limit }){
    state.limit = limit
  },
  SET_CURRENT_PAGE(state, { currentPage }){
    state.currentPage = (currentPage - 1)
  },
  SET_LOADING(state){
    state.loading = !state.loading
  },
  SET_ICAT(state, {icat}){
    if(icat){
      state.cat = icat
    }
  },
  SET_SAL(state, {salary}){
    state.salarym = salary
  }
}

export default mutations
