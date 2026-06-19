'use client';

import { useState } from 'react';

export default function Home() {
    const [count, setCount] = useState<string>('');
    const [luckyNumbers, setLuckyNumbers] = useState<string[]>([]);
    const [error, setError] = useState<string>('');
    const [loading, setLoading] = useState<boolean>(false);

    const handleRandom = async (e: React.FormEvent) => {
        e.preventDefault();
        setError('');

        const num = parseInt(count);
        if (isNaN(num) || num < 1 || num > 10) {
            setError('กรุณาใส่ตัวเลข 1 - 10 เท่านั้น');
            return;
        }

        setLoading(true);

        try {
            const res = await fetch('http://localhost:8080/api/random-numbers', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ count: num }),
            });

            const result = await res.json();

            if (res.ok) {
                setLuckyNumbers(result.data);
            } else {
                setError(result.error || 'เกิดข้อผิดพลาดในการเชื่อมต่อ');
            }
        } catch (err) {
            setError('ไม่สามารถเชื่อมต่อกับ Server หลังบ้านได้');
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="flex justify-center items-center min-h-screen bg-gray-50 p-4">
            {/* การ์ดหลักภายนอก */}
            <div className="w-full max-w-md bg-white border-2 border-black rounded-[32px] p-8 flex flex-col items-center shadow-sm">

                <h1 className="text-xl font-bold mb-2 text-black">สุ่มเบอร์มงคล</h1>
                <label className="text-sm mb-2 text-black">ใส่ตัวเลข</label>

                {/* ฟอร์มกรอกตัวเลข */}
                <form onSubmit={handleRandom} className="text-black w-full flex flex-col items-center mb-6">
                    <input
                        type="number"
                        min="1"
                        max="10"
                        value={count}
                        onChange={(e) => setCount(e.target.value)}
                        placeholder="1-10"
                        className="w-32 border-2 border-black rounded-xl px-3 py-1 text-center font-medium focus:outline-none"
                    />
                    <button
                        type="submit"
                        disabled={loading}
                        className="mt-3 px-4 py-1.5 bg-black text-white rounded-lg text-sm hover:bg-gray-800 transition"
                    >
                        {loading ? 'กำลังสุ่ม...' : 'สุ่มเบอร์'}
                    </button>
                </form>

                {error && <p className="text-red-500 text-xs mb-4 text-center">{error}</p>}

                {/* กล่องแสดงผลลัพธ์ (List Box) */}
                <div className="w-full min-h-[250px] border-2 border-black rounded-[24px] overflow-hidden bg-white">
                    {luckyNumbers.length > 0 ? (
                        <div className="divide-y-2 divide-black">
                            {luckyNumbers.map((num, index) => (
                                <div
                                    key={index}
                                    className="py-3 text-center text-black font-mono text-lg font-semibold tracking-wider bg-white"
                                >
                                    {num}
                                </div>
                            ))}
                        </div>
                    ) : (
                        <div className="h-[250px] flex items-center justify-center text-black text-sm p-4 text-center">
                            ผลลัพธ์เบอร์มงคลจะแสดงที่นี่
                        </div>
                    )}
                </div>

            </div>
        </div>
    );
}