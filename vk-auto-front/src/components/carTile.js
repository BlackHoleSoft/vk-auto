export default function CarTile({ car }) {
    return (
        <div className="car-tile">
            <div className="car-tile--name">{car.Brand} {car.Model}</div>
            <div className="car-tile--price">{car.Price}₽</div>
            <div className="car-tile--year">{car.Year} год</div>
            <div className="car-tile--desc">{car.Description}</div>
            <div className="car-tile--photos">
                {car.Photos.map((el, i) => (
                    <img key={i} src={el} width="200px" alt=""/>
                ))}
            </div>
        </div>
    )
}