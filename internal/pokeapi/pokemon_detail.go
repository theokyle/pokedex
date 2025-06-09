package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) PokemonDetail(pokemon *string) (RespPokemon, error) {
	url := baseURL + "/pokemon/" + *pokemon

	data, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespPokemon{}, err
		}
		res, err := c.httpClient.Do(req)
		if err != nil {
			return RespPokemon{}, err
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return RespPokemon{}, err
		}

		c.cache.Add(url, data)
	}

	pokemonResp := RespPokemon{}
	if err := json.Unmarshal(data, &pokemonResp); err != nil {
		return RespPokemon{}, err
	}

	return pokemonResp, nil
}
