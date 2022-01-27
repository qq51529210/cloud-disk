import axios from "axios";

const client = axios.create({
  baseURL: "http://127.0.0.1:3390",
  timeout: 3000,
});

const errCodeParseJSON = 0;
const errCodeFormValue = 1;
const errCodeQueryData = 2;
const errCodeUnauthorized = 3;

const errorReason = code => {
  switch (code) {
    case errCodeParseJSON:
      return "JSON数据错误";
    case errCodeFormValue:
      return "URL参数错误";
    case errCodeQueryData:
      return "查询数据出错";
    case errCodeUnauthorized:
      return "账号或密码错误";
    default:
      return "错误的服务器响应";
  }
};

const onReject = rej => {
  return {
    error:
      rej.response && rej.response.data.error
        ? errorReason(rej.response.data.code)
        : "请求出错",
  };
};

const post = (url, model) => {
  return client
    .post(url, model, {
      validateStatus: status => status === 201,
    })
    .then(res => {
      return {
        data: res.data,
      };
    })
    .catch(onReject);
};

export const signInAccount = model => post("/api/tokens?type=account", model);

export const signInPhone = model => post("/api/tokens?type=phone", model);

export const signUp = model => post("/users", model);

export const getPhoneCode = number => post("/");
