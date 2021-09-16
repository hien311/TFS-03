import { FETCH_ARTICLE, FETCH_ARTICLES, FETCH_COMMENTS, FETCH_USER_ARTICLES, POST_ARTICLES } from "./actions.type";
import ArticlesService from "../common/articles.service";
import { SET_ERRORS, SET_ARTICLES, SET_ARTICLES_COUNT, SET_RESPONSE, SET_ARTICLE, SET_COMMENTS, SET_AUTHOR, SET_USER_ARTICLES } from "./mutations.type";

const state = {
  articles: [],
  articlesCount: 0,
  userArticles: "",
  article: "",
  author: "",
  comments: "",
  postResponse: "",
  tags: [],
  errors: null,
};

const getters = {};

const actions = {
  async [FETCH_ARTICLE]({ commit }, slug) {
    try {
      const { data } = await ArticlesService.get(slug)
      const { article} = data
      commit(SET_ARTICLE, article)
      commit(SET_AUTHOR, article.author)
    } catch(err) {
      commit(SET_ERRORS, err.response)
    }
  },
  async [FETCH_ARTICLES]({ commit }, { type, ...payload }) {
    try {
      const { data } = await ArticlesService.query(type, payload);
      const { articles, articlesCount } = data;
      commit(SET_ARTICLES, articles);
      commit(SET_ARTICLES_COUNT, articlesCount);
    } catch (err) {
      commit(SET_ERRORS, err.response.data.errors);
    }
  },
  async [FETCH_COMMENTS]({ commit }, slug) {
    try {
      const { data } = await ArticlesService.get(slug)
      const { comments } = data 
      commit(SET_COMMENTS, comments)
    } catch(err) {
      commit(SET_ERRORS, err.response)
    }
  }
  ,
  async [POST_ARTICLES]( { commit }, payload) {
    try {
      const response =  await ArticlesService.create(payload)
      commit(SET_RESPONSE, response.data)
      return true
    } catch(err) {
      commit(SET_ERRORS, err.response)
      return false
    }
  },
  async [FETCH_USER_ARTICLES]({ commit }, payload) {
    try {
      const { data } = await ArticlesService.query("", payload)
      const { articles } = data
      commit(SET_USER_ARTICLES, articles)
    } catch(err) {
      commit(SET_ERRORS, err.response)
    }
  }
};

const mutations = {
  [SET_ARTICLES](state, articles) {
    state.articles = articles;
  },
  [SET_ARTICLES_COUNT](state, articlesCount) {
    state.articlesCount = articlesCount;
  },
  [SET_RESPONSE](state, response) {
    state.postResponse = response
  },
  [SET_ARTICLE](state, article){
    state.article = article
  },
  [SET_COMMENTS](state, comments) {
    state.comments = comments
  },
  [SET_AUTHOR](state, author){
    state.author = author
  },
  [SET_USER_ARTICLES](state, articles) {
    state.userArticles = articles
  },
};

export default {
  state,
  getters,
  actions,
  mutations,
};
