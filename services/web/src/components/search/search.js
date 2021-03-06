import { motion } from "framer-motion";
import React, { useEffect, useRef, useState } from "react";

import useDebounce from "../../hooks/debounce";

import SearchResultList from "./resultList";

function Search(query) {
    const queryString = `q=${query}`;
    return fetch(`http://localhost:8080/api/search?${queryString}`, {
        mode: "cors"
    })
        .then((r) => {
            return r.json();
        })
        .catch((error) => {
            console.log(error);
            return [];
        });
}

const SearchInput = () => {
    const [query, setQuery] = useState("");
    const [results, setResults] = useState([]);
    const [isSearching, setIsSearching] = useState(false);
    const [inputWidth, setInputWidth] = useState(0);
    const [focused, setFocused] = useState(false);

    const debouncedQuery = useDebounce(query, 200);

    useEffect(() => {
        if (debouncedQuery) {
            setIsSearching(true);
            Search(query).then((results) => {
                setIsSearching(false);
                setResults(results);
            });
        } else {
            setResults([]);
        }
    }, [debouncedQuery]);

    const inputRef = useRef(null);
    useEffect(() => {
        setInputWidth(inputRef.current ? inputRef.current?.offsetWidth : 0);
    }, [inputRef.current?.offsetWidth]);

    return (
        <div className="relative h-full mx-4">
            <input
                type="text"
                autoComplete="off"
                ref={inputRef}
                className="relative font-sans md:py-0 py-2 bg-white w-full h-full text-black outline-none px-4 rounded"
                id="global-search-input"
                placeholder="Search"
                onChange={(e) => setQuery(e.target.value)}
                onBlur={() => {
                    setFocused(false);
                }}
                onFocus={() => {
                    setFocused(true);
                }}
            />
            {query && focused && (
                <div>
                    <hr className="md:hidden bg-black h-1px opacity-50" />
                    <div
                        className="md:popup"
                        style={{
                            width: inputWidth
                        }}>
                        {
                            //isSearching && <span className="text-black">Searching...</span>
                        }
                        <SearchResultList results={results} />
                    </div>
                </div>
            )}
        </div>
    );
};

export default SearchInput;
