package data

import "fmt"

type Project struct {
    Title       string
    Stack       []string
    Description string
    URL         string
}

func GetWorkData(p string) (Project, error) {
    projects := map[string]Project{
        "one-with-the-riverbed": {
            "One With the Riverbed",
            []string{"NextJS", "TypeScript", "Tailwind", "Vercel"},
            "One With the Riverbed is a band based out of Kalamazoo, MI.",
            "https://onewiththeriverbed.com",
        },
        "aarons-music-service": {
            "Aaron's Music Service",
            []string{"Svelte", "SCSS", "Render"},
            "Aaron's Music Service is a luthier and consignment shop in Vicksburg, MI.",
            "https://aaronsmusicservice.com",
        },
        "get-pokemon-data-from-api": {
            "Get Pokemon Data From API",
            []string{"TypeScript", "Jest", "Rest API", "Render"},
            "This is the first step to creating a full-stack pokemon app. This is how I got the data I needed to create the database.",
            "https://github.com/BradPreston/getPokemonDataFromAPI",
        },
        "kazoo48": {
            "Kazoo48 Film Festival",
            []string{"NextJS", "TypeScript", "SCSS", "Stripe", "Vercel"},
            "Kazoo48 is a film festival in Kalamazoo, MI that allows teams to compete to create a film in 48 hours",
            "https://kazoo48.com",
        },
    }

    var r Project

    result, ok := projects[p]

    if !ok {
        return r, fmt.Errorf("%s", "no results found")
    }

    r = result

    return r, nil
}

