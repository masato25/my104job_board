const mutations = {
  SET_JOBS_LIST(state, { list }){
    state.jobs = list.jobs
    state.total = list.count
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
    state.cat = icat
  },
  SET_SAL(state, {salary}){
    state.salarym = salary
  }
}

export default mutations
