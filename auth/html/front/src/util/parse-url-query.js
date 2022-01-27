export const parseUrlQuery = queryStr => {
  if (queryStr === "") {
    return {};
  }
  if (queryStr[0] === "?") {
    queryStr = queryStr.slice(1);
  }
  let data = {};
  queryStr.split("&").forEach(v => {
    let i = v.indexOf("=");
    if (i > 0) {
      data[v.slice(0, i)] = v.slice(i + 1);
    }
  });
  return data;
};
