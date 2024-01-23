'use client';
import React, { useEffect, useState } from 'react';
import Image from 'next/image';

export default function Home() {
    const [currentQuestion, setCurrentQuestion] = useState([]);
    const [isFinalQuestion, setFinalQuestion] = useState(false);
    const [showFinalResults, setFinalResults] = useState(false);
    const [showFinalResultScreen, setFinalResultScreen] = useState(false);
    const [displayQuestions, setDisplayQuestions] = useState(false);
    const [displayLandingPage, setDisplayLandingPage] = useState(true);
    const [totalQuestions, setTotalQuestions] = useState(0);
    const [score, setScore] = useState(0);
    const [answer, setAnswer] = useState('');
    const [hasNotAnswered, setHasNotAnswered] = useState(false);
    const [currID, setCurrID] = useState(1);
    const [difficulty, setDifficulty] = useState('easy');
    const [titleDiff, setTitleDiff] = useState('');
    const [colorDiff, setColorDiff] = useState('');
    const [playerRank, setPlayerRank] = useState('');
    const [playerRankedImage, setPlayerRankedImage] = useState('');

    const BASE_URL = 'https://mighty-cove-41770-527d6e6e093b.herokuapp.com';

    useEffect(() => {
        fetch(`${BASE_URL}/questions/${difficulty}/${currID}`)
            .then((data) => data.json())
            .then((data) => {
                setCurrentQuestion(data);
            });
    }, [difficulty, currID]);

    const getTotalQuestions = (diff) => {
        fetch(`${BASE_URL}/totalquestions/${diff}`)
            .then((data) => data.json())
            .then((data) => {
                setTotalQuestions(data);
            });
    };

    const diffSetting = (difficulty) => {
        setDifficulty(difficulty);
        getTotalQuestions(difficulty);
        if (difficulty === 'easy') {
            setTitleDiff('Easy');
            setColorDiff('green');
        } else if (difficulty === 'medium') {
            setTitleDiff('Medium');
            setColorDiff('orange');
        } else if (difficulty === 'hard') {
            setTitleDiff('Hard');
            setColorDiff('red');
        }

        setCurrID(1);
        setDisplayLandingPage(false);
        setDisplayQuestions(true);
    };

    const optionClicked = (question) => {
        setAnswer(question);
        setHasNotAnswered(false);
    };

    const nextQuestion = (currentID, diff) => {
        if (answer === '') {
            verifyAnswer();
            return;
        }

        getScore(currentID, answer, diff, score);
        setHasNotAnswered(false);
        setAnswer('');
        currentID++;
        setCurrID(currentID);

        if (currentID === totalQuestions) {
            setFinalQuestion(true);
            return;
        }
    };

    const verifyAnswer = () => {
        setHasNotAnswered(true);
    };

    const getScore = (currentID, answer, diff, currScore) => {
        fetch(`${BASE_URL}/questions/score`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                id: currentID,
                answer: answer,
                diff: diff,
                currScore: currScore,
            }),
        })
            .then((data) => data.json())
            .then((data) => {
                setScore(data);
            });
    };

    const getPlayerRank = () => {
        setFinalResults(true);

        if ((score / totalQuestions) * 100 >= 90) {
            setPlayerRank('a Master');
            setPlayerRankedImage('/images/t_rankBigIcon_06.png');
        } else if ((score / totalQuestions) * 100 >= 80) {
            setPlayerRank('an Ultra');
            setPlayerRankedImage('/images/t_rankBigIcon_05.png');
        } else if ((score / totalQuestions) * 100 >= 70) {
            setPlayerRank('a Veteran');
            setPlayerRankedImage('/images/t_rankBigIcon_04.png');
        } else if ((score / totalQuestions) * 100 >= 60) {
            setPlayerRank('an Expert');
            setPlayerRankedImage('/images/t_rankBigIcon_03.png');
        } else if ((score / totalQuestions) * 100 >= 50) {
            setPlayerRank('a Great');
            setPlayerRankedImage('/images/t_rankBigIcon_02.png');
        } else if ((score / totalQuestions) * 100 <= 49) {
            setPlayerRank('a Beginner');
            setPlayerRankedImage('/images/t_rankBigIcon_01.png');
        }
    };

    const seeResults = (currentID, diff) => {
        if (answer === '') {
            verifyAnswer();
            return;
        }

        setHasNotAnswered(false);
        getScore(currentID, answer, diff, score);
        setDisplayQuestions(false);
        setFinalResultScreen(true);
    };

    const restartQuiz = () => {
        setFinalResultScreen(false);
        setFinalResults(false);
        setFinalQuestion(false);
        setDisplayLandingPage(true);
        setDifficulty('');
        setScore(0);
        setCurrID(1);
        setAnswer('');
        setPlayerRank('');
        setPlayerRankedImage('');
    };

    return (
        <div className="app">
            <Image
                src="/images/bgimggr.jpg"
                alt="Image from The Pokemon Company/Tencent Games/"
                fill
                sizes="100vw"
                style={{
                    objectFit: 'cover',
                    zIndex: -2,
                }}
            />
            {displayLandingPage ? (
                <div className="landing-page">
                    <h2>Pokemon UNITE Quiz</h2>
                    <h3>Welcome!</h3>
                    <h4>
                        This is a quiz based on Pokémon UNITE, a Free-to-Play,
                        Multiplayer Online Battle Arena Video Game, developed by
                        TiMi Studio Group and published by The Pokémon Company
                        for Android and iOS, and by Nintendo for the Nintendo
                        Switch. If you would like to know more, you can visit
                        the{' '}
                        <a href="https://unite.pokemon.com/en-us/">
                            Official Pokemon UNITE Webpage
                        </a>
                        .
                    </h4>
                    <h4>
                        Below you can select a difficulty level based on your
                        experience with UNITE. <br /> At the end, you will be
                        given a UNITE Ranked Trophy based on your performance!
                    </h4>
                    <h3>Select a difficulty to begin.</h3>
                    <button
                        className="easy-button"
                        onClick={() => diffSetting('easy')}
                    >
                        Easy
                    </button>
                    <button
                        className="medium-button"
                        onClick={() => diffSetting('medium')}
                    >
                        Medium
                    </button>
                    <button
                        className="hard-button"
                        onClick={() => diffSetting('hard')}
                    >
                        Hard
                    </button>
                </div>
            ) : (
                <div> </div>
            )}

            {displayQuestions ? (
                <React.Fragment>
                    <div className="diff-card">
                        <h2 className={colorDiff}>{titleDiff} Questions</h2>
                    </div>
                    <div key={currentQuestion.id} className="question-card">
                        <h3 className="question-text">
                            Question {currentQuestion.id}:{' '}
                            {currentQuestion.title}
                        </h3>
                        <h3 className="question-text">Your Answer: {answer}</h3>
                        {hasNotAnswered ? (
                            <h3 className="hasnt-answered-text">
                                You haven't selected an answer.
                            </h3>
                        ) : (
                            <h3></h3>
                        )}
                        <ul className="option-holder">
                            <li
                                onClick={() => optionClicked('A')}
                                className="option"
                            >
                                A: {currentQuestion.a}
                            </li>
                            <li
                                onClick={() => optionClicked('B')}
                                className="option"
                            >
                                B: {currentQuestion.b}
                            </li>
                            <li
                                onClick={() => optionClicked('C')}
                                className="option"
                            >
                                C: {currentQuestion.c}
                            </li>
                            <li
                                onClick={() => optionClicked('D')}
                                className="option"
                            >
                                D: {currentQuestion.d}
                            </li>
                        </ul>
                    </div>
                    {isFinalQuestion ? (
                        <div className="see-result-card">
                            <h1>See Results</h1>
                            <button
                                className="restart-button"
                                onClick={() =>
                                    seeResults(
                                        currentQuestion.id,
                                        currentQuestion.diff
                                    )
                                }
                            >
                                Click Here!
                            </button>
                        </div>
                    ) : (
                        <div className="next-card">
                            <button
                                className="next-button"
                                onClick={() =>
                                    nextQuestion(
                                        currentQuestion.id,
                                        currentQuestion.diff
                                    )
                                }
                            >
                                Next Question
                            </button>
                        </div>
                    )}
                </React.Fragment>
            ) : (
                <div></div>
            )}

            {showFinalResultScreen ? (
                <div className="final-result">
                    <h1> Final Results </h1>

                    {showFinalResults ? (
                        <div>
                            <h2>
                                {score} out of {totalQuestions} correct - (
                                {(score / totalQuestions) * 100}%)
                            </h2>
                            <div>
                                <Image
                                    src={playerRankedImage}
                                    alt="Trophy Rank image from The Pokemon Company/Tencent Games"
                                    width={111}
                                    height={206}
                                    sizes="30vw"
                                    style={{
                                        objectFit: 'cover',
                                        zIndex: -2,
                                    }}
                                />
                                <h1>You are {playerRank} Ranked Player!</h1>
                            </div>
                            <button
                                className="restart-button"
                                onClick={() => restartQuiz()}
                            >
                                Restart
                            </button>
                        </div>
                    ) : (
                        <div>
                            <button
                                className="get-finalscore-button"
                                onClick={() => getPlayerRank()}
                            >
                                {' '}
                                Click here to show your final score!
                            </button>
                        </div>
                    )}
                </div>
            ) : (
                <div></div>
            )}
            <div className="bottom-space"></div>
            <footer className="footer">
                <h3 className="footer-text">
                    Pokémon UNITE and all images used are owned by © 2023
                    Pokémon. ©1995–2023 Nintendo Creatures Inc. / GAME FREAK
                    inc. © 2023 Tencent. All rights reserved.
                </h3>
            </footer>
        </div>
    );
}
