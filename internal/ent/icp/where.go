// Code generated by ent, DO NOT EDIT.

package icp

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/yoshino-s/soar-helper/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Icp {
	return predicate.Icp(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Icp {
	return predicate.Icp(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Icp {
	return predicate.Icp(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Icp {
	return predicate.Icp(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Icp {
	return predicate.Icp(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Icp {
	return predicate.Icp(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Icp {
	return predicate.Icp(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Icp {
	return predicate.Icp(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Icp {
	return predicate.Icp(sql.FieldLTE(FieldID, id))
}

// Host applies equality check predicate on the "host" field. It's identical to HostEQ.
func Host(v string) predicate.Icp {
	return predicate.Icp(sql.FieldEQ(FieldHost, v))
}

// City applies equality check predicate on the "city" field. It's identical to CityEQ.
func City(v string) predicate.Icp {
	return predicate.Icp(sql.FieldEQ(FieldCity, v))
}

// Province applies equality check predicate on the "province" field. It's identical to ProvinceEQ.
func Province(v string) predicate.Icp {
	return predicate.Icp(sql.FieldEQ(FieldProvince, v))
}

// Company applies equality check predicate on the "company" field. It's identical to CompanyEQ.
func Company(v string) predicate.Icp {
	return predicate.Icp(sql.FieldEQ(FieldCompany, v))
}

// Owner applies equality check predicate on the "owner" field. It's identical to OwnerEQ.
func Owner(v string) predicate.Icp {
	return predicate.Icp(sql.FieldEQ(FieldOwner, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Icp {
	return predicate.Icp(sql.FieldEQ(FieldType, v))
}

// Homepage applies equality check predicate on the "homepage" field. It's identical to HomepageEQ.
func Homepage(v string) predicate.Icp {
	return predicate.Icp(sql.FieldEQ(FieldHomepage, v))
}

// Permit applies equality check predicate on the "permit" field. It's identical to PermitEQ.
func Permit(v string) predicate.Icp {
	return predicate.Icp(sql.FieldEQ(FieldPermit, v))
}

// WebName applies equality check predicate on the "webName" field. It's identical to WebNameEQ.
func WebName(v string) predicate.Icp {
	return predicate.Icp(sql.FieldEQ(FieldWebName, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Icp {
	return predicate.Icp(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Icp {
	return predicate.Icp(sql.FieldEQ(FieldUpdatedAt, v))
}

// HostEQ applies the EQ predicate on the "host" field.
func HostEQ(v string) predicate.Icp {
	return predicate.Icp(sql.FieldEQ(FieldHost, v))
}

// HostNEQ applies the NEQ predicate on the "host" field.
func HostNEQ(v string) predicate.Icp {
	return predicate.Icp(sql.FieldNEQ(FieldHost, v))
}

// HostIn applies the In predicate on the "host" field.
func HostIn(vs ...string) predicate.Icp {
	return predicate.Icp(sql.FieldIn(FieldHost, vs...))
}

// HostNotIn applies the NotIn predicate on the "host" field.
func HostNotIn(vs ...string) predicate.Icp {
	return predicate.Icp(sql.FieldNotIn(FieldHost, vs...))
}

// HostGT applies the GT predicate on the "host" field.
func HostGT(v string) predicate.Icp {
	return predicate.Icp(sql.FieldGT(FieldHost, v))
}

// HostGTE applies the GTE predicate on the "host" field.
func HostGTE(v string) predicate.Icp {
	return predicate.Icp(sql.FieldGTE(FieldHost, v))
}

// HostLT applies the LT predicate on the "host" field.
func HostLT(v string) predicate.Icp {
	return predicate.Icp(sql.FieldLT(FieldHost, v))
}

// HostLTE applies the LTE predicate on the "host" field.
func HostLTE(v string) predicate.Icp {
	return predicate.Icp(sql.FieldLTE(FieldHost, v))
}

// HostContains applies the Contains predicate on the "host" field.
func HostContains(v string) predicate.Icp {
	return predicate.Icp(sql.FieldContains(FieldHost, v))
}

// HostHasPrefix applies the HasPrefix predicate on the "host" field.
func HostHasPrefix(v string) predicate.Icp {
	return predicate.Icp(sql.FieldHasPrefix(FieldHost, v))
}

// HostHasSuffix applies the HasSuffix predicate on the "host" field.
func HostHasSuffix(v string) predicate.Icp {
	return predicate.Icp(sql.FieldHasSuffix(FieldHost, v))
}

// HostEqualFold applies the EqualFold predicate on the "host" field.
func HostEqualFold(v string) predicate.Icp {
	return predicate.Icp(sql.FieldEqualFold(FieldHost, v))
}

// HostContainsFold applies the ContainsFold predicate on the "host" field.
func HostContainsFold(v string) predicate.Icp {
	return predicate.Icp(sql.FieldContainsFold(FieldHost, v))
}

// CityEQ applies the EQ predicate on the "city" field.
func CityEQ(v string) predicate.Icp {
	return predicate.Icp(sql.FieldEQ(FieldCity, v))
}

// CityNEQ applies the NEQ predicate on the "city" field.
func CityNEQ(v string) predicate.Icp {
	return predicate.Icp(sql.FieldNEQ(FieldCity, v))
}

// CityIn applies the In predicate on the "city" field.
func CityIn(vs ...string) predicate.Icp {
	return predicate.Icp(sql.FieldIn(FieldCity, vs...))
}

// CityNotIn applies the NotIn predicate on the "city" field.
func CityNotIn(vs ...string) predicate.Icp {
	return predicate.Icp(sql.FieldNotIn(FieldCity, vs...))
}

// CityGT applies the GT predicate on the "city" field.
func CityGT(v string) predicate.Icp {
	return predicate.Icp(sql.FieldGT(FieldCity, v))
}

// CityGTE applies the GTE predicate on the "city" field.
func CityGTE(v string) predicate.Icp {
	return predicate.Icp(sql.FieldGTE(FieldCity, v))
}

// CityLT applies the LT predicate on the "city" field.
func CityLT(v string) predicate.Icp {
	return predicate.Icp(sql.FieldLT(FieldCity, v))
}

// CityLTE applies the LTE predicate on the "city" field.
func CityLTE(v string) predicate.Icp {
	return predicate.Icp(sql.FieldLTE(FieldCity, v))
}

// CityContains applies the Contains predicate on the "city" field.
func CityContains(v string) predicate.Icp {
	return predicate.Icp(sql.FieldContains(FieldCity, v))
}

// CityHasPrefix applies the HasPrefix predicate on the "city" field.
func CityHasPrefix(v string) predicate.Icp {
	return predicate.Icp(sql.FieldHasPrefix(FieldCity, v))
}

// CityHasSuffix applies the HasSuffix predicate on the "city" field.
func CityHasSuffix(v string) predicate.Icp {
	return predicate.Icp(sql.FieldHasSuffix(FieldCity, v))
}

// CityIsNil applies the IsNil predicate on the "city" field.
func CityIsNil() predicate.Icp {
	return predicate.Icp(sql.FieldIsNull(FieldCity))
}

// CityNotNil applies the NotNil predicate on the "city" field.
func CityNotNil() predicate.Icp {
	return predicate.Icp(sql.FieldNotNull(FieldCity))
}

// CityEqualFold applies the EqualFold predicate on the "city" field.
func CityEqualFold(v string) predicate.Icp {
	return predicate.Icp(sql.FieldEqualFold(FieldCity, v))
}

// CityContainsFold applies the ContainsFold predicate on the "city" field.
func CityContainsFold(v string) predicate.Icp {
	return predicate.Icp(sql.FieldContainsFold(FieldCity, v))
}

// ProvinceEQ applies the EQ predicate on the "province" field.
func ProvinceEQ(v string) predicate.Icp {
	return predicate.Icp(sql.FieldEQ(FieldProvince, v))
}

// ProvinceNEQ applies the NEQ predicate on the "province" field.
func ProvinceNEQ(v string) predicate.Icp {
	return predicate.Icp(sql.FieldNEQ(FieldProvince, v))
}

// ProvinceIn applies the In predicate on the "province" field.
func ProvinceIn(vs ...string) predicate.Icp {
	return predicate.Icp(sql.FieldIn(FieldProvince, vs...))
}

// ProvinceNotIn applies the NotIn predicate on the "province" field.
func ProvinceNotIn(vs ...string) predicate.Icp {
	return predicate.Icp(sql.FieldNotIn(FieldProvince, vs...))
}

// ProvinceGT applies the GT predicate on the "province" field.
func ProvinceGT(v string) predicate.Icp {
	return predicate.Icp(sql.FieldGT(FieldProvince, v))
}

// ProvinceGTE applies the GTE predicate on the "province" field.
func ProvinceGTE(v string) predicate.Icp {
	return predicate.Icp(sql.FieldGTE(FieldProvince, v))
}

// ProvinceLT applies the LT predicate on the "province" field.
func ProvinceLT(v string) predicate.Icp {
	return predicate.Icp(sql.FieldLT(FieldProvince, v))
}

// ProvinceLTE applies the LTE predicate on the "province" field.
func ProvinceLTE(v string) predicate.Icp {
	return predicate.Icp(sql.FieldLTE(FieldProvince, v))
}

// ProvinceContains applies the Contains predicate on the "province" field.
func ProvinceContains(v string) predicate.Icp {
	return predicate.Icp(sql.FieldContains(FieldProvince, v))
}

// ProvinceHasPrefix applies the HasPrefix predicate on the "province" field.
func ProvinceHasPrefix(v string) predicate.Icp {
	return predicate.Icp(sql.FieldHasPrefix(FieldProvince, v))
}

// ProvinceHasSuffix applies the HasSuffix predicate on the "province" field.
func ProvinceHasSuffix(v string) predicate.Icp {
	return predicate.Icp(sql.FieldHasSuffix(FieldProvince, v))
}

// ProvinceIsNil applies the IsNil predicate on the "province" field.
func ProvinceIsNil() predicate.Icp {
	return predicate.Icp(sql.FieldIsNull(FieldProvince))
}

// ProvinceNotNil applies the NotNil predicate on the "province" field.
func ProvinceNotNil() predicate.Icp {
	return predicate.Icp(sql.FieldNotNull(FieldProvince))
}

// ProvinceEqualFold applies the EqualFold predicate on the "province" field.
func ProvinceEqualFold(v string) predicate.Icp {
	return predicate.Icp(sql.FieldEqualFold(FieldProvince, v))
}

// ProvinceContainsFold applies the ContainsFold predicate on the "province" field.
func ProvinceContainsFold(v string) predicate.Icp {
	return predicate.Icp(sql.FieldContainsFold(FieldProvince, v))
}

// CompanyEQ applies the EQ predicate on the "company" field.
func CompanyEQ(v string) predicate.Icp {
	return predicate.Icp(sql.FieldEQ(FieldCompany, v))
}

// CompanyNEQ applies the NEQ predicate on the "company" field.
func CompanyNEQ(v string) predicate.Icp {
	return predicate.Icp(sql.FieldNEQ(FieldCompany, v))
}

// CompanyIn applies the In predicate on the "company" field.
func CompanyIn(vs ...string) predicate.Icp {
	return predicate.Icp(sql.FieldIn(FieldCompany, vs...))
}

// CompanyNotIn applies the NotIn predicate on the "company" field.
func CompanyNotIn(vs ...string) predicate.Icp {
	return predicate.Icp(sql.FieldNotIn(FieldCompany, vs...))
}

// CompanyGT applies the GT predicate on the "company" field.
func CompanyGT(v string) predicate.Icp {
	return predicate.Icp(sql.FieldGT(FieldCompany, v))
}

// CompanyGTE applies the GTE predicate on the "company" field.
func CompanyGTE(v string) predicate.Icp {
	return predicate.Icp(sql.FieldGTE(FieldCompany, v))
}

// CompanyLT applies the LT predicate on the "company" field.
func CompanyLT(v string) predicate.Icp {
	return predicate.Icp(sql.FieldLT(FieldCompany, v))
}

// CompanyLTE applies the LTE predicate on the "company" field.
func CompanyLTE(v string) predicate.Icp {
	return predicate.Icp(sql.FieldLTE(FieldCompany, v))
}

// CompanyContains applies the Contains predicate on the "company" field.
func CompanyContains(v string) predicate.Icp {
	return predicate.Icp(sql.FieldContains(FieldCompany, v))
}

// CompanyHasPrefix applies the HasPrefix predicate on the "company" field.
func CompanyHasPrefix(v string) predicate.Icp {
	return predicate.Icp(sql.FieldHasPrefix(FieldCompany, v))
}

// CompanyHasSuffix applies the HasSuffix predicate on the "company" field.
func CompanyHasSuffix(v string) predicate.Icp {
	return predicate.Icp(sql.FieldHasSuffix(FieldCompany, v))
}

// CompanyIsNil applies the IsNil predicate on the "company" field.
func CompanyIsNil() predicate.Icp {
	return predicate.Icp(sql.FieldIsNull(FieldCompany))
}

// CompanyNotNil applies the NotNil predicate on the "company" field.
func CompanyNotNil() predicate.Icp {
	return predicate.Icp(sql.FieldNotNull(FieldCompany))
}

// CompanyEqualFold applies the EqualFold predicate on the "company" field.
func CompanyEqualFold(v string) predicate.Icp {
	return predicate.Icp(sql.FieldEqualFold(FieldCompany, v))
}

// CompanyContainsFold applies the ContainsFold predicate on the "company" field.
func CompanyContainsFold(v string) predicate.Icp {
	return predicate.Icp(sql.FieldContainsFold(FieldCompany, v))
}

// OwnerEQ applies the EQ predicate on the "owner" field.
func OwnerEQ(v string) predicate.Icp {
	return predicate.Icp(sql.FieldEQ(FieldOwner, v))
}

// OwnerNEQ applies the NEQ predicate on the "owner" field.
func OwnerNEQ(v string) predicate.Icp {
	return predicate.Icp(sql.FieldNEQ(FieldOwner, v))
}

// OwnerIn applies the In predicate on the "owner" field.
func OwnerIn(vs ...string) predicate.Icp {
	return predicate.Icp(sql.FieldIn(FieldOwner, vs...))
}

// OwnerNotIn applies the NotIn predicate on the "owner" field.
func OwnerNotIn(vs ...string) predicate.Icp {
	return predicate.Icp(sql.FieldNotIn(FieldOwner, vs...))
}

// OwnerGT applies the GT predicate on the "owner" field.
func OwnerGT(v string) predicate.Icp {
	return predicate.Icp(sql.FieldGT(FieldOwner, v))
}

// OwnerGTE applies the GTE predicate on the "owner" field.
func OwnerGTE(v string) predicate.Icp {
	return predicate.Icp(sql.FieldGTE(FieldOwner, v))
}

// OwnerLT applies the LT predicate on the "owner" field.
func OwnerLT(v string) predicate.Icp {
	return predicate.Icp(sql.FieldLT(FieldOwner, v))
}

// OwnerLTE applies the LTE predicate on the "owner" field.
func OwnerLTE(v string) predicate.Icp {
	return predicate.Icp(sql.FieldLTE(FieldOwner, v))
}

// OwnerContains applies the Contains predicate on the "owner" field.
func OwnerContains(v string) predicate.Icp {
	return predicate.Icp(sql.FieldContains(FieldOwner, v))
}

// OwnerHasPrefix applies the HasPrefix predicate on the "owner" field.
func OwnerHasPrefix(v string) predicate.Icp {
	return predicate.Icp(sql.FieldHasPrefix(FieldOwner, v))
}

// OwnerHasSuffix applies the HasSuffix predicate on the "owner" field.
func OwnerHasSuffix(v string) predicate.Icp {
	return predicate.Icp(sql.FieldHasSuffix(FieldOwner, v))
}

// OwnerIsNil applies the IsNil predicate on the "owner" field.
func OwnerIsNil() predicate.Icp {
	return predicate.Icp(sql.FieldIsNull(FieldOwner))
}

// OwnerNotNil applies the NotNil predicate on the "owner" field.
func OwnerNotNil() predicate.Icp {
	return predicate.Icp(sql.FieldNotNull(FieldOwner))
}

// OwnerEqualFold applies the EqualFold predicate on the "owner" field.
func OwnerEqualFold(v string) predicate.Icp {
	return predicate.Icp(sql.FieldEqualFold(FieldOwner, v))
}

// OwnerContainsFold applies the ContainsFold predicate on the "owner" field.
func OwnerContainsFold(v string) predicate.Icp {
	return predicate.Icp(sql.FieldContainsFold(FieldOwner, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Icp {
	return predicate.Icp(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Icp {
	return predicate.Icp(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Icp {
	return predicate.Icp(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Icp {
	return predicate.Icp(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Icp {
	return predicate.Icp(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Icp {
	return predicate.Icp(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Icp {
	return predicate.Icp(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Icp {
	return predicate.Icp(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Icp {
	return predicate.Icp(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Icp {
	return predicate.Icp(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Icp {
	return predicate.Icp(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Icp {
	return predicate.Icp(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Icp {
	return predicate.Icp(sql.FieldContainsFold(FieldType, v))
}

// HomepageEQ applies the EQ predicate on the "homepage" field.
func HomepageEQ(v string) predicate.Icp {
	return predicate.Icp(sql.FieldEQ(FieldHomepage, v))
}

// HomepageNEQ applies the NEQ predicate on the "homepage" field.
func HomepageNEQ(v string) predicate.Icp {
	return predicate.Icp(sql.FieldNEQ(FieldHomepage, v))
}

// HomepageIn applies the In predicate on the "homepage" field.
func HomepageIn(vs ...string) predicate.Icp {
	return predicate.Icp(sql.FieldIn(FieldHomepage, vs...))
}

// HomepageNotIn applies the NotIn predicate on the "homepage" field.
func HomepageNotIn(vs ...string) predicate.Icp {
	return predicate.Icp(sql.FieldNotIn(FieldHomepage, vs...))
}

// HomepageGT applies the GT predicate on the "homepage" field.
func HomepageGT(v string) predicate.Icp {
	return predicate.Icp(sql.FieldGT(FieldHomepage, v))
}

// HomepageGTE applies the GTE predicate on the "homepage" field.
func HomepageGTE(v string) predicate.Icp {
	return predicate.Icp(sql.FieldGTE(FieldHomepage, v))
}

// HomepageLT applies the LT predicate on the "homepage" field.
func HomepageLT(v string) predicate.Icp {
	return predicate.Icp(sql.FieldLT(FieldHomepage, v))
}

// HomepageLTE applies the LTE predicate on the "homepage" field.
func HomepageLTE(v string) predicate.Icp {
	return predicate.Icp(sql.FieldLTE(FieldHomepage, v))
}

// HomepageContains applies the Contains predicate on the "homepage" field.
func HomepageContains(v string) predicate.Icp {
	return predicate.Icp(sql.FieldContains(FieldHomepage, v))
}

// HomepageHasPrefix applies the HasPrefix predicate on the "homepage" field.
func HomepageHasPrefix(v string) predicate.Icp {
	return predicate.Icp(sql.FieldHasPrefix(FieldHomepage, v))
}

// HomepageHasSuffix applies the HasSuffix predicate on the "homepage" field.
func HomepageHasSuffix(v string) predicate.Icp {
	return predicate.Icp(sql.FieldHasSuffix(FieldHomepage, v))
}

// HomepageIsNil applies the IsNil predicate on the "homepage" field.
func HomepageIsNil() predicate.Icp {
	return predicate.Icp(sql.FieldIsNull(FieldHomepage))
}

// HomepageNotNil applies the NotNil predicate on the "homepage" field.
func HomepageNotNil() predicate.Icp {
	return predicate.Icp(sql.FieldNotNull(FieldHomepage))
}

// HomepageEqualFold applies the EqualFold predicate on the "homepage" field.
func HomepageEqualFold(v string) predicate.Icp {
	return predicate.Icp(sql.FieldEqualFold(FieldHomepage, v))
}

// HomepageContainsFold applies the ContainsFold predicate on the "homepage" field.
func HomepageContainsFold(v string) predicate.Icp {
	return predicate.Icp(sql.FieldContainsFold(FieldHomepage, v))
}

// PermitEQ applies the EQ predicate on the "permit" field.
func PermitEQ(v string) predicate.Icp {
	return predicate.Icp(sql.FieldEQ(FieldPermit, v))
}

// PermitNEQ applies the NEQ predicate on the "permit" field.
func PermitNEQ(v string) predicate.Icp {
	return predicate.Icp(sql.FieldNEQ(FieldPermit, v))
}

// PermitIn applies the In predicate on the "permit" field.
func PermitIn(vs ...string) predicate.Icp {
	return predicate.Icp(sql.FieldIn(FieldPermit, vs...))
}

// PermitNotIn applies the NotIn predicate on the "permit" field.
func PermitNotIn(vs ...string) predicate.Icp {
	return predicate.Icp(sql.FieldNotIn(FieldPermit, vs...))
}

// PermitGT applies the GT predicate on the "permit" field.
func PermitGT(v string) predicate.Icp {
	return predicate.Icp(sql.FieldGT(FieldPermit, v))
}

// PermitGTE applies the GTE predicate on the "permit" field.
func PermitGTE(v string) predicate.Icp {
	return predicate.Icp(sql.FieldGTE(FieldPermit, v))
}

// PermitLT applies the LT predicate on the "permit" field.
func PermitLT(v string) predicate.Icp {
	return predicate.Icp(sql.FieldLT(FieldPermit, v))
}

// PermitLTE applies the LTE predicate on the "permit" field.
func PermitLTE(v string) predicate.Icp {
	return predicate.Icp(sql.FieldLTE(FieldPermit, v))
}

// PermitContains applies the Contains predicate on the "permit" field.
func PermitContains(v string) predicate.Icp {
	return predicate.Icp(sql.FieldContains(FieldPermit, v))
}

// PermitHasPrefix applies the HasPrefix predicate on the "permit" field.
func PermitHasPrefix(v string) predicate.Icp {
	return predicate.Icp(sql.FieldHasPrefix(FieldPermit, v))
}

// PermitHasSuffix applies the HasSuffix predicate on the "permit" field.
func PermitHasSuffix(v string) predicate.Icp {
	return predicate.Icp(sql.FieldHasSuffix(FieldPermit, v))
}

// PermitIsNil applies the IsNil predicate on the "permit" field.
func PermitIsNil() predicate.Icp {
	return predicate.Icp(sql.FieldIsNull(FieldPermit))
}

// PermitNotNil applies the NotNil predicate on the "permit" field.
func PermitNotNil() predicate.Icp {
	return predicate.Icp(sql.FieldNotNull(FieldPermit))
}

// PermitEqualFold applies the EqualFold predicate on the "permit" field.
func PermitEqualFold(v string) predicate.Icp {
	return predicate.Icp(sql.FieldEqualFold(FieldPermit, v))
}

// PermitContainsFold applies the ContainsFold predicate on the "permit" field.
func PermitContainsFold(v string) predicate.Icp {
	return predicate.Icp(sql.FieldContainsFold(FieldPermit, v))
}

// WebNameEQ applies the EQ predicate on the "webName" field.
func WebNameEQ(v string) predicate.Icp {
	return predicate.Icp(sql.FieldEQ(FieldWebName, v))
}

// WebNameNEQ applies the NEQ predicate on the "webName" field.
func WebNameNEQ(v string) predicate.Icp {
	return predicate.Icp(sql.FieldNEQ(FieldWebName, v))
}

// WebNameIn applies the In predicate on the "webName" field.
func WebNameIn(vs ...string) predicate.Icp {
	return predicate.Icp(sql.FieldIn(FieldWebName, vs...))
}

// WebNameNotIn applies the NotIn predicate on the "webName" field.
func WebNameNotIn(vs ...string) predicate.Icp {
	return predicate.Icp(sql.FieldNotIn(FieldWebName, vs...))
}

// WebNameGT applies the GT predicate on the "webName" field.
func WebNameGT(v string) predicate.Icp {
	return predicate.Icp(sql.FieldGT(FieldWebName, v))
}

// WebNameGTE applies the GTE predicate on the "webName" field.
func WebNameGTE(v string) predicate.Icp {
	return predicate.Icp(sql.FieldGTE(FieldWebName, v))
}

// WebNameLT applies the LT predicate on the "webName" field.
func WebNameLT(v string) predicate.Icp {
	return predicate.Icp(sql.FieldLT(FieldWebName, v))
}

// WebNameLTE applies the LTE predicate on the "webName" field.
func WebNameLTE(v string) predicate.Icp {
	return predicate.Icp(sql.FieldLTE(FieldWebName, v))
}

// WebNameContains applies the Contains predicate on the "webName" field.
func WebNameContains(v string) predicate.Icp {
	return predicate.Icp(sql.FieldContains(FieldWebName, v))
}

// WebNameHasPrefix applies the HasPrefix predicate on the "webName" field.
func WebNameHasPrefix(v string) predicate.Icp {
	return predicate.Icp(sql.FieldHasPrefix(FieldWebName, v))
}

// WebNameHasSuffix applies the HasSuffix predicate on the "webName" field.
func WebNameHasSuffix(v string) predicate.Icp {
	return predicate.Icp(sql.FieldHasSuffix(FieldWebName, v))
}

// WebNameIsNil applies the IsNil predicate on the "webName" field.
func WebNameIsNil() predicate.Icp {
	return predicate.Icp(sql.FieldIsNull(FieldWebName))
}

// WebNameNotNil applies the NotNil predicate on the "webName" field.
func WebNameNotNil() predicate.Icp {
	return predicate.Icp(sql.FieldNotNull(FieldWebName))
}

// WebNameEqualFold applies the EqualFold predicate on the "webName" field.
func WebNameEqualFold(v string) predicate.Icp {
	return predicate.Icp(sql.FieldEqualFold(FieldWebName, v))
}

// WebNameContainsFold applies the ContainsFold predicate on the "webName" field.
func WebNameContainsFold(v string) predicate.Icp {
	return predicate.Icp(sql.FieldContainsFold(FieldWebName, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Icp {
	return predicate.Icp(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Icp {
	return predicate.Icp(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Icp {
	return predicate.Icp(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Icp {
	return predicate.Icp(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Icp {
	return predicate.Icp(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Icp {
	return predicate.Icp(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Icp {
	return predicate.Icp(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Icp {
	return predicate.Icp(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Icp {
	return predicate.Icp(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Icp {
	return predicate.Icp(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Icp {
	return predicate.Icp(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Icp {
	return predicate.Icp(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Icp {
	return predicate.Icp(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Icp {
	return predicate.Icp(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Icp {
	return predicate.Icp(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Icp {
	return predicate.Icp(sql.FieldLTE(FieldUpdatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Icp) predicate.Icp {
	return predicate.Icp(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Icp) predicate.Icp {
	return predicate.Icp(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Icp) predicate.Icp {
	return predicate.Icp(sql.NotPredicates(p))
}
