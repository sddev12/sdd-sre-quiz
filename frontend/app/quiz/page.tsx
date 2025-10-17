"use client";
import React from "react";
import { QuizComponent } from "./QuizComponent";
import { useSearchParams } from "next/navigation";

// Dummy implementations for now
const fetchQuestion = async (index: number) => {
    const questions = [
        {
            id: "q1",
            text: "What does SRE stand for?",
            answers: [
                { id: "a1", text: "Site Reliability Engineer" },
                { id: "a2", text: "Software Reliability Engineer" },
                { id: "a3", text: "Systems Reliability Engineer" },
                { id: "a4", text: "Secure Reliability Engineer" },
            ],
        },
        // ...more questions
    ];
    return questions[index % questions.length];
};

const submitAnswer = async (questionId: string, answerId: string) => {
    // Dummy: only a1 is correct
    return { correct: answerId === "a1" };
};

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
                    submitAnswer={submitAnswer}
                />
            </main>
            <footer className="bg-white border-t-2 border-black py-2 px-4 flex justify-between text-xs font-mono">
                <span>Version 0.1</span>
                <span>© 2025 SRE QUIZ</span>
            </footer>
        </div>
    );
};
