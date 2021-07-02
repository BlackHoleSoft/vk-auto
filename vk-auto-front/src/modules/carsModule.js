import React, { useEffect, useState } from 'react'
import CarTile from '../components/carTile';

export default function CarsModule() {
    let [cars, setCars] = useState([]);

    useEffect(() => {
        fetch("/json/cars.json")
            .then(response => response.json())
            .then(json => {
                console.log(json);
                setCars(json);
            });
    }, [])

    return cars
        .sort((a, b) => a.Price - b.Price)
        .map(el => <CarTile key={el.Id} car={el} />);
}