function minEatingSpeed(piles, h) {
        let result = 1;

        while (true) {
            let time = 0
            for (const pile of piles) {
                time += Math.ceil(pile / result)
            }

            if (time <= h) {
                break;
            }

            result++
        }

        return result
    }

    console.log(minEatingSpeed([25,10,23,4], 4))