import axios from 'axios';

const actions = {
    getJobs({ commit, state }, params){
      let cat = state.cat
      let limit = state.limit
      let salarym = state.salarym
      if (params != undefined){
        cat =  params["cat"] || state.cat
        limit =  params["limit"] || state.limit
        salarym =  params["salary"] || state.salarym
      }
      commit('SET_LOADING')
      axios.get(`/api/jobs?cat=${cat}&limit=${limit}&page=${state.currentPage}&salary=${salarym}`).then((response) => {
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
    setIcat({ commit, state }, {icat}){
      commit('SET_ICAT', {icat})
    },
    setSalary({commit}, {salary}){
      commit('SET_SAL', {salary})
    }
}

export default actions
