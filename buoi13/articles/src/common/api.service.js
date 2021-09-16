import axios from "axios";
import API_URL from "./config";
import JwtService from "./jwt.service";

axios.defaults.baseURL = API_URL;

const ApiService = {
  setHeader() {
    axios.defaults.headers.common[
      "Authorization"
    ] = `Token ${JwtService.getToken()}`;
  },

  query(resource, params) {
    return axios.get(resource, params);
  },

  get(resource, slug = "") {
    return axios.get(`${resource}/${slug}`);
  },

  post(resource, params) {
    return axios.post(`${resource}`, params);
  },

  update(resource, slug, params) {
    return axios.put(`${resource}/${slug}`, params);
  },

  put(resource, params) {
    return axios.put(`${resource}`, params);
  },

  delete(resource) {
    return axios.delete(resource);
  },
};

export default ApiService;
