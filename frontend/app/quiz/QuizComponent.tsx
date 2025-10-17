"use client";
import React, { useState, useEffect } from "react";

export interface QuizComponentProps {
    userName: string;
    totalQuestions: number;
    onQuizComplete: (result: { score: number; time: number }) => void;
    fetchQuestion: (index: number) => Promise<{
        id: string;
        text: string;
        answers: { id: string; text: string }[];
    }>;
    submitAnswer: (questionId: string, answerId: string) => Promise<{ correct: boolean }>;
    initialTime?: number;
}

export const QuizComponent: React.FC<QuizComponentProps> = ({
    userName,
    totalQuestions,
    onQuizComplete,
    fetchQuestion,
    submitAnswer,
    initialTime = 0,
}) => {
    const [currentQuestionIndex, setCurrentQuestionIndex] = useState(0);
    const [currentQuestion, setCurrentQuestion] = useState<{
        id: string;
        text: string;
        answers: { id: string; text: string }[];
    } | null>(null);
    const [selectedAnswerId, setSelectedAnswerId] = useState<string | null>(null);
    const [elapsedTime, setElapsedTime] = useState(initialTime);
    const [score, setScore] = useState(0);
    const [isComplete, setIsComplete] = useState(false);
    const [isLoading, setIsLoading] = useState(false);

    // Timer
    useEffect(() => {
        if (isComplete) return;
        const timer = setInterval(() => setElapsedTime((t) => t + 1), 1000);
        return () => clearInterval(timer);
    }, [isComplete]);

    // Fetch question
    useEffect(() => {
        setIsLoading(true);
        fetchQuestion(currentQuestionIndex)
            .then((q) => setCurrentQuestion(q))
            .finally(() => setIsLoading(false));
        setSelectedAnswerId(null);
    }, [currentQuestionIndex, fetchQuestion]);

    const handleSelect = (answerId: string) => {
        setSelectedAnswerId(answerId);
    };

    const handleSubmit = async () => {
        if (!currentQuestion || !selectedAnswerId) return;
        setIsLoading(true);
        const res = await submitAnswer(currentQuestion.id, selectedAnswerId);
        if (res.correct) setScore((s) => s + 1);
        if (currentQuestionIndex + 1 < totalQuestions) {
            setCurrentQuestionIndex((i) => i + 1);
        } else {
            setIsComplete(true);
            onQuizComplete({ score: res.correct ? score + 1 : score, time: elapsedTime });
        }
        setIsLoading(false);
    };

    // Format time MM:SS
    const formatTime = (s: number) => `${String(Math.floor(s / 60)).padStart(2, "0")}:${String(s % 60).padStart(2, "0")}`;

    return (
        <div className="max-w-xl mx-auto border-4 border-black bg-white p-8 mt-10 shadow-lg">
            <div className="flex justify-between mb-4 text-xs font-mono uppercase text-black">
                <span>Question {isComplete ? totalQuestions : currentQuestionIndex + 1} of {totalQuestions}</span>
                <span>Time: {formatTime(elapsedTime)}</span>
            </div>
            <div className="mb-2 text-sm font-mono text-black">User: <span className="font-bold">{userName}</span></div>
            {!isComplete && currentQuestion && (
                <>
                    <div className="text-2xl font-bold uppercase text-center mb-6 font-mono text-black">{currentQuestion.text}</div>
                    <div className="flex flex-col gap-4 mb-6">
                        {currentQuestion.answers.map((a) => (
                            <button
                                key={a.id}
                                className={`w-full border-2 rounded py-3 px-4 text-lg font-mono uppercase transition-all ${selectedAnswerId === a.id ? "bg-black text-white border-black" : "bg-white text-black border-black hover:bg-gray-200"}`}
                                onClick={() => handleSelect(a.id)}
                                disabled={isLoading}
                            >
                                {a.text}
                            </button>
                        ))}
                    </div>
                    <button
                        className="w-full border-2 border-black bg-white text-black font-bold uppercase py-3 rounded disabled:opacity-50"
                        onClick={handleSubmit}
                        disabled={!selectedAnswerId || isLoading}
                    >
                        Submit Answer
                    </button>
                </>
            )}
            {isComplete && (
                <div className="text-center mt-8">
                    <div className="text-2xl font-bold uppercase mb-4 text-black">Quiz Complete!</div>
                    <div className="mb-2 font-mono text-black">Score: <span className="font-bold">{score}/{totalQuestions}</span></div>
                    <div className="mb-6 font-mono text-black">Time: {formatTime(elapsedTime)}</div>
                    <button className="border-2 border-black bg-white text-black font-bold uppercase py-3 px-6 rounded hover:bg-black hover:text-white transition-all">
                        View Leaderboard
                    </button>
                </div>
            )}
        </div>
    );
};
