import axios from 'axios';

const actions = {
    getJobs({ commit, state }){
      let cat = state.cat
      let limit = state.limit
      let salarym = state.salarym
      commit('SET_LOADING')
      axios.get(`/api/jobs?cat=${cat}&per_page=${limit}&page=${state.currentPage}&salary=${salarym}`).then((response) => {
         commit('SET_JOBS_LIST', { list: response.data })
         commit('SET_LOADING')
       }, (err) => {
         console.log(err)
       })
    },
    setCurrentPage({ commit, dispatch }, { page }){
      commit('SET_CURRENT_PAGE', {currentPage: page})
      dispatch('getJobs')
    },
    setPageShowLimit({ commit, dispatch }, {limit}){
      commit('SET_PAGE_Limit', {limit})
      dispatch('getJobs')
    },
    setIcat({ commit, dispatch}, {icat}){
      commit('SET_ICAT', {icat})
      dispatch('getJobs')
    },
    setSalary({commit, dispatch}, {salary}){
      commit('SET_SAL', {salary})
      dispatch('getJobs')
    }
}

export default actions
