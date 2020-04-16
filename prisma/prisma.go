package prisma

import (
	//"Tiamat/log"
	"Tiamat/prisma/prisma-client"
	"context"
)

var (
	Prisclient  *prisma.Client
	Priscontext context.Context
)

func InitPrisma() {
	Prisclient = prisma.New(nil)
	Priscontext = context.TODO()

	return
}

func Login(UserName string, Password string) (string, int) {
	user, err := Prisclient.User(prisma.UserWhereUniqueInput{
		UserName: &UserName,
	}).Exec(Priscontext)
	if err != nil || user.PassWord != Password {
		return "nil", 0
	}
	return user.ID, 1
}

func GetIp(Domain string) string {
	ip, err := Prisclient.Domain(prisma.DomainWhereUniqueInput{
		Name: &Domain,
	}).Exec(Priscontext)
	if err != nil {
		return "Denied"
	}
	return ip.Ip
}

func Add(Domain string, IP string, User string) string {
	_, err := Prisclient.CreateDomain(prisma.DomainCreateInput{
		Name:   Domain,
		Ip:     IP,
		Author: User,
	}).Exec(Priscontext)
	if err != nil {
		return "Denied"
	}
	return "Accept"
}

func Remove(Domain string, User string) string {
	check, err := Prisclient.DeleteManyDomains(&prisma.DomainWhereInput{
		And: []prisma.DomainWhereInput{
			{
				Name: &Domain,
			},
			{
				Author: &User,
			},
		},
	}).Exec(Priscontext)
	if err != nil || check.Count == 0 {
		return "Denied"
	}
	return "Accept"
}

func Update(Domain string, IP string, User string) string {
	check, err := Prisclient.UpdateManyDomains(prisma.DomainUpdateManyParams{
		Where: &prisma.DomainWhereInput{
			And: []prisma.DomainWhereInput{
				{
					Name: &Domain,
				},
				{
					Author: &User,
				},
			},
		},
		Data: prisma.DomainUpdateManyMutationInput{
			Ip: &IP,
		},
	}).Exec(Priscontext)
	if err != nil || check.Count == 0 {
		return "Denied"
	}
	return "Accept"
}
