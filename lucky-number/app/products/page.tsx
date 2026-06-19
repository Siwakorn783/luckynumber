// import React from 'react'
// type Product = {
//     id: number;
//     name: string;
//     price: number;
// }
// async function getProducts(): Promise<Product[]> {
//     return [
//         { id: 1, name: "Basic NextJs", price: 1000 },
//         { id: 2, name: "Basic reactJs", price: 1000 },
//         { id: 3, name: "Basic Js", price: 1000 },
//     ]

// }
// async function ProductsPage() {

//     const products = await getProducts();
//     return (
//         <main className='p-8'>
//             <h1 className='text-3xl font-bold'> รายการสินค้า</h1>
//             <div className='mt-6 grid gap-4'>
//                 {products.map((product) => (
//                     <div key={product.id} className='rounded-lg border p-4'>
//                         <h2 className='text-xl font-semibold'>{product.name}</h2>
//                         <p className='mt-2'>ราคา {product.price} บาท</p>
//                     </div>
//                 ))}

//             </div>

//         </main>
//     )
// }

// export default ProductsPage