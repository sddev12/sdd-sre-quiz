"use client";
import React from "react";
import { QuizComponent } from "./QuizComponent";
import { useSearchParams } from "next/navigation";


// Fetch question from backend API
const fetchQuestion = async (index: number) => {
    // For now, always fetch question 01 for index 0
    const questionId = String(index + 1).padStart(2, '0');
    const res = await fetch(`http://localhost:8080/question?questionId=${questionId}`);
    if (!res.ok) {
        throw new Error('Failed to fetch question');
    }
    const data = await res.json();
    // Map backend response to QuizComponent format
    type Answer = { answerId: string; answer: string };
    return {
        id: data.questionId,
        text: data.question,
        answers: (data.answers || []).map((a: Answer) => ({ id: a.answerId, text: a.answer })),
    };
};

const submitAnswer = async (questionId: string, answerId: string, userName?: string) => {
    // Use the provided userName from props or fallback
    const username = userName || '';
    const res = await fetch('http://localhost:8080/submit-answer', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ username, questionId, answerId }),
    });
    if (res.ok) {
        return { correct: true };
    } else {
        const data = await res.json().catch(() => ({}));
        throw new Error(data.message || 'Failed to submit answer');
    }
};


// Wrapper to inject userName from QuizPage
function makeSubmitAnswerWithUserName(userName: string) {
    return (questionId: string, answerId: string) => submitAnswer(questionId, answerId, userName);
}
export default function QuizPage() {
    const searchParams = useSearchParams();
    const userName = searchParams.get("name") || "";
    const handleQuizComplete = (result: { score: number; time: number }) => {
        // TODO: handle completion (e.g., navigate to leaderboard)
        alert(`Quiz complete! Score: ${result.score}, Time: ${result.time}s`);
    };

    return (
        <div className="min-h-screen bg-gray-100 flex flex-col">
            <header className="bg-black text-white text-xl font-bold uppercase py-4 text-center tracking-widest">
                Quiz Terminal
            </header>
            <main className="flex-1 flex flex-col justify-center">
                <QuizComponent
                    userName={userName}
                    totalQuestions={20}
                    onQuizComplete={handleQuizComplete}
                    fetchQuestion={fetchQuestion}
                    submitAnswer={makeSubmitAnswerWithUserName(userName)}
                />
            </main>
            <footer className="bg-white border-t-2 border-black py-2 px-4 flex justify-between text-xs font-mono">
                <span>Version 0.1</span>
                <span>© 2025 SRE QUIZ</span>
            </footer>
        </div>
    );
};
