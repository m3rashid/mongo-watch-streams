const mongoose = require('mongoose');
const { faker } = require('@faker-js/faker');

mongoose.set('debug', true);

const productSchema = new mongoose.Schema(
	{
		name: String,
		price: Number,
		stock: Number,
	},
	{ timestamps: true }
);
const Product = mongoose.model('Product', productSchema);

const userSchema = new mongoose.Schema(
	{
		username: String,
		password: String,
		score: Number,
		address: String,
	},
	{ timestamps: false }
);
const User = mongoose.model('User', userSchema);

const addProduct = async () => {
	const productData = {
		name: faker.commerce.productName(),
		price: faker.commerce.price(),
		stock: faker.number.int(10000),
	};

	const product = new Product(productData);
	await product.save();
};

const addUser = async () => {
	const userData = {
		username: faker.person.fullName(),
		password: faker.internet.password(),
		score: faker.number.int(100),
		address: faker.location.streetAddress(),
	};

	const user = new User(userData);
	await user.save();
};

const sleep = (ms) => {
	return new Promise((resolve) => setTimeout(resolve, ms));
};

const main = async () => {
	await mongoose.connect('<db_uri>', { useNewUrlParser: true, useUnifiedTopology: true, });
	console.log('Database connected');

	for (let i = 0; i < 10; i++) {
		await addProduct();
		await sleep(1000);

		await addUser();
		await sleep(1000);
	}

	mongoose.disconnect();
	console.log('Database disconnected');
	return;
};

main().catch(console.error);
