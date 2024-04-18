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

export const localDate = (inp) => {
	let ns = new Date(inp)
	return ns.toLocaleString()
}

export const USDollar = new Intl.NumberFormat('en-US', {
	style: 'currency',
	currency: 'USD',
	minimumFractionDigits: 8,
});