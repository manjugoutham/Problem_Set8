/* Address Book of MyFriends: Create an Address-Book of your friends. It will have fields like - name, Mobile Number, Alternate Mobile number,
Address, City/Town Name. Some friends will have Alternate Mobile numbers and some will not have. Populate the same with 10 of your friend's details.
Write this in a file. After every friend detail there should 2 blank lines in the file  */

package main

import (
	"fmt"
	"os"
)

type Friend struct {
	Name            string
	MobileNumber    string
	AlternateNumber string
	Address         string
	CityOrTownName  string
}

func main() {
	friends := []Friend{
		{
			Name:            "Goutham",
			MobileNumber:    "9998876543",
			AlternateNumber: "9987654321",
			Address:         "banglore",
			CityOrTownName:  "newyark",
		},
		{
			Name:            "Manju",
			MobileNumber:    "9008876543",
			AlternateNumber: "9980654321",
			Address:         "manglore",
			CityOrTownName:  "Goa",
		},
		{
			Name:            "Chandu",
			MobileNumber:    "9876543211",
			AlternateNumber: "9876543221",
			Address:         "kgf",
			CityOrTownName:  "kolar",
		},
		{
			Name:            "Abhi",
			MobileNumber:    "9876543111",
			AlternateNumber: "9876543121",
			Address:         "banglore",
			CityOrTownName:  "kgf",
		},
	}

	file, err := os.Create("friends.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	for _, friend := range friends {
		_, err := fmt.Fprintf(file, "Name: %s\nMobile Number: %s\n", friend.Name, friend.MobileNumber)
		if err != nil {
			fmt.Println(err)
		}

		if friend.AlternateNumber != "" {
			fmt.Fprintf(file, "Alternate Mobile Number: %s\n", friend.AlternateNumber)
		}

		fmt.Fprintf(file, "Address: %s\nCity/Town Name: %s\n\n\n", friend.Address, friend.CityOrTownName)
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("successful")
}
