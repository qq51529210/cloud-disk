import axios from "axios";
import { baseURL, error, onError } from "./api";

export interface model {
    account: string;
    password: string;
}

// post 用于登录请求
export const post = async (m: model): Promise<error | null> => {
    // form data
    let data = new FormData()
    data.append('account', m.account)
    data.append('password', m.password)
    // 请求
    return axios.postForm(baseURL('/login'), data, {
        validateStatus: status => status == 200
    }).then(res => {
        return null
    }).catch(err => {
        return onError(err)
    })
}