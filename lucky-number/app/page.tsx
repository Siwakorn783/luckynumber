"use client";
import { useState } from 'react';

export default function Home() {
  // state สำหรับเก็บเบอร์
  const [phone, setPhone] = useState('');

  // ฟังก์ชันนี้เตรียมไว้ให้ Go Backend เรียกใช้ในอนาคต
  const handleRequestRandomFromGo = async () => {
    try {
      // เพิ่มตำแหน่งนี้ลงไปครับ:
      const response = await fetch('http://localhost:8080/api/random-number', {
        cache: 'no-store' // <--- ตรงนี้คือจุดที่ต้องเพิ่มครับ
      });




      const data = await response.json();
      setPhone(data.phoneNumber);
    } catch (error) {
      alert("ยังไม่ได้รับ Go Backend หรือ API ยังไม่พร้อม");
    }
  };

  return (
    <main className="p-8 max-w-sm mx-auto">
      <h1 className="text-xl font-bold mb-4 text-center">ระบบเบอร์มงคล</h1>

      <div className="flex flex-col gap-4">
        {/* ช่องสุ่มเบอร์ (ปุ่มที่ไปขอเลขจาก Go) */}
        <button
          className="bg-green-600 text-black p-2 rounded"
          onClick={handleRequestRandomFromGo}
        >
          ขอเลขสุ่ม
        </button>

        {/* ช่องใส่เลข */}
        <input
          type="text" // ใช้เป็น text เพื่อให้แสดงขีดหรือรูปแบบเบอร์ได้สวยงาม
          className="border p-2 rounded w-full bg-gray-100 text-black cursor-not-allowed" // ใส่สีเทาเพื่อให้ดูเป็นช่องแสดงผล
          placeholder="กดปุ่มเพื่อรับเลขมงคล"
          value={phone}
          readOnly // เพิ่มคำสั่งนี้เพื่อห้ามผู้ใช้พิมพ์แก้ไข
        />

      </div>
    </main>
  );
}
