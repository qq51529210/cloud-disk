import axios from "axios";
import { baseURL, formContentType } from "./api";

export interface model {
    account: string;
    password: string;
}

// post 用于登录请求
export const post = async (m: model): Promise<Boolean> => {
    // form data
    let data = new FormData()
    data.append('account', m.account)
    data.append('password', m.password)
    // url
    let url = baseURL + '/login'
    if (window.location.search) {
        url += window.location.search
    }
    // 请求
    return axios.postForm(url, data).
        then(res => {
            return true
        }).catch(err => {
            return false
        })
}