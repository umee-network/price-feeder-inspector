export const getReq = async (url) => {
	return await fetch(url)
		.then((resp) => resp.json())
		.then((resp) => {
			return resp;
		})
		.catch((err) => {
			console.log(err);
		});
};