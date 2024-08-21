"use client"

import React, {useEffect, useState} from 'react';
import axios from 'axios';
import './page.css'

interface Kols {
	kolID: number;
	userProfileID: number;
	language: string;
	education: string;
	expectedSalary: number;
	expectedSalaryEnable: boolean;
	channelSettingTypeID: number;
	idFrontURL: string;
	idBackURL: string;
	portraitURL: string;
	rewardID: number;
	paymentMethodID: number;
	testimonialsID: number;
	verificationStatus: boolean;
	enabled: boolean;
	activeDate: Date;
	active: boolean;
	createdBy: string;
	createdDate: Date;
	modifiedBy: string;
	modifiedDate: Date;
	isRemove: boolean;
	isOnBoarding: boolean;
	code: string;
	portraitRightURL: string;
	portraitLeftURL: string;
	livenessStatus: boolean;
}

const Page = () => {
    // * Use useState to store Kols from API 
    // ! (if you have more optimized way to store data, PLEASE FEELS FREE TO CHANGE)
    const [Kols , setKols] = useState<Kols[]>([]);  
	const [loading, setLoading] = useState(true);
    // * Fetch API over here 
    // * Use useEffect to fetch data from API 
	const [currentPage, setCurrentPage] = useState(1);
	const pageLimit = 5;
	const [searchLanguage, setSearchLanguage] = useState('');

    useEffect(() => {
		const fetchKols = async () => {
			setLoading(true);
			try {
				const response = await axios.get('https://testing-api-ch6p.onrender.com/kols', {
					params: {
						pageIndex: 1,
						pageLimit: 30
					}
				});
				if (Array.isArray(response.data.kol)) {
                    setKols(response.data.kol);
                }
				else {
                    console.error('Expected an array but got:', response.data.kol);
                }
			} catch (error) {
				console.error('Error fetching data:', error);
			} finally {
				setLoading(false);
            }
		};
		fetchKols();
    }, []);

	const fetchKols_searchParams = async () => {
		setLoading(true);
		try {
			const searchParams = JSON.stringify(
				{"key":"language","value": searchLanguage}
			);
			// searchParams={"key":"language","value":"English"}
			const response = await axios.get('https://testing-api-ch6p.onrender.com/kols', {
				params: {
					pageIndex: 1,
					pageLimit: 30,
					searchParams: searchParams
				}
			});
			if (Array.isArray(response.data.kol)) {
				setKols(response.data.kol);
			}
			else {
				console.error('Expected an array but got:', response.data.kol);
			}
		} catch (error) {
			console.error('Error fetching data:', error);
		} finally {
			setLoading(false);
		}
	};

	const scrollToTop = () => {
		if (currentPage > 1) {
			triggerScrollAnimation("up");
			setTimeout(() => {
			setCurrentPage(currentPage - 1);
			}, 500); // Match this to the animation duration
		}
	};

	const scrollToBottom = () => {
		if (currentPage * pageLimit < Kols.length) {
			triggerScrollAnimation("down");
			setTimeout(() => {
			setCurrentPage(currentPage + 1);
			}, 500); // Match this to the animation duration
		}
	};

	const triggerScrollAnimation = (direction: "up" | "down") => {
		const tableElement = document.querySelector(".kols-table tbody") as HTMLElement;
		if (tableElement) {
			tableElement.classList.remove("scroll-up", "scroll-down");
			// Force reflow to restart animation
			void tableElement.offsetWidth;
			tableElement.classList.add(direction === "up" ? "scroll-up" : "scroll-down");
		}
	};

	const handleWheel = (event: WheelEvent) => {
		event.preventDefault();
		if (event.deltaY < 0) {
			scrollToTop();
		} else if (event.deltaY > 0) {
			scrollToBottom();
		}
	};

	useEffect(() => {
		const tableElement = document.querySelector(".kols-table") as HTMLElement;
		if (tableElement) {
			tableElement.addEventListener("wheel", handleWheel, { passive: false });
		}
		return () => {
			if (tableElement) {
				tableElement.removeEventListener("wheel", handleWheel);
			}
		};
	}, [currentPage]);

	const displayedKols = Kols.slice((currentPage - 1) * pageLimit, currentPage * pageLimit);
    const totalPages = Math.ceil(Kols.length / pageLimit);
    return (
        <div className="container">
            <h1 className="header">Kols List</h1>
			<div className="search-form">
				<input
					type="text"
					placeholder="Search by language"
					value={searchLanguage}
					onChange={(e) => setSearchLanguage(e.target.value)}
					onKeyPress={(e) => {
						if (e.key === 'Enter') {
							setCurrentPage(1); // Reset to the first page when searching
							fetchKols_searchParams(); // Fetch KOLs with new search parameters
						}
					}}
				/>
				<button onClick={() => {
					setCurrentPage(1); // Reset to the first page when searching
					fetchKols_searchParams(); // Fetch KOLs with new search parameters
				}}>Search</button>
			</div>
            <div className="scroll-buttons">
                <button onClick={scrollToTop} className="scroll-button" disabled={currentPage === 1}>Scroll Up</button>
                <button onClick={scrollToBottom} className="scroll-button" disabled={(currentPage * pageLimit) >= Kols.length}>Scroll Down</button>
            </div>
            {loading ? (
                <p>Loading...</p>
            ) : (
                <>
                    <div className="d-flex">
                        {Kols.length === 0 ? (
                            <p>No KOLs available.</p>
                        ) : (
                            <table className="kols-table">
                                <thead>
                                    <tr>
                                        <th>KolID</th>
                                        <th>UserProfileID</th>
                                        <th>Language</th>
                                        <th>Education</th>
                                        <th>Expected Salary</th>
										<th>Expected Salary Enable</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {displayedKols.map((kol, index) => (
                                        <tr key={`${kol.kolID}-${index}`}>
                                            <td>{kol.kolID}</td>
                                            <td>{kol.userProfileID}</td>
                                            <td>{kol.language}</td>
                                            <td>{kol.education}</td>
                                            <td>{kol.expectedSalary}</td>
											<td>{kol.expectedSalaryEnable.toString()}</td>
                                        </tr>
                                    ))}
                                </tbody>
                            </table>
                        )}
                    </div>
                    <div className="page-indicator">
                        Page {currentPage} of {totalPages}
                    </div>
                </>
            )}
        </div>
    );
};

export default Page;