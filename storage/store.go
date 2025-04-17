package storage

import "structure-phonebook-cli/models"

var Contacts []models.Contact

func AddContact(contact models.Contact) {
	Contacts = append(Contacts, contact)

}
func GetAllContacts() []models.Contact {
	return Contacts
}

func FindContactByName(name string) *models.Contact {
	for i, contact := range Contacts {
		if contact.Name == name {
			return &Contacts[i]
		}
	}
	return nil
}
